package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/go-wonk/si"
	"github.com/go-wonk/si/sigorm"
	"github.com/go-wonk/si/sihttp"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	commonadapter "github.com/w-woong/common/adapter"
	"github.com/w-woong/common/configs"
	"github.com/w-woong/common/logger"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/common/wrapper"
	"github.com/w-woong/order/adapter"
	"github.com/w-woong/order/cmd/route"
	"github.com/w-woong/order/port"
	"github.com/w-woong/order/usecase"
	"gorm.io/gorm"
)

var (
	Version = "undefined"

	printVersion     bool
	tickIntervalSec  int = 30
	addr             string
	certPem, certKey string
	readTimeout      int
	writeTimeout     int
	configName       string
	maxProc          int

	usePprof  = false
	pprofAddr = ":56060"
)

func init() {
	flag.StringVar(&addr, "addr", ":49003", "listen address")
	flag.BoolVar(&printVersion, "version", false, "print version")
	flag.IntVar(&tickIntervalSec, "tick", 30, "tick interval in second")
	flag.StringVar(&certKey, "key", "./certs/key.pem", "server key")
	flag.StringVar(&certPem, "pem", "./certs/cert.pem", "server pem")
	flag.IntVar(&readTimeout, "readTimeout", 30, "read timeout")
	flag.IntVar(&writeTimeout, "writeTimeout", 30, "write timeout")
	flag.StringVar(&configName, "config", "./configs/server.yml", "config file name")
	flag.IntVar(&maxProc, "mp", runtime.NumCPU(), "GOMAXPROCS")

	flag.BoolVar(&usePprof, "pprof", false, "use pprof")
	flag.StringVar(&pprofAddr, "pprof_addr", ":56060", "pprof listen address")

	flag.Parse()
}

func main() {
	// defaultTimeout := 6 * time.Second

	var err error

	if printVersion {
		fmt.Printf("version \"%v\"\n", Version)
		return
	}
	runtime.GOMAXPROCS(maxProc)

	// config
	conf := common.Config{}
	if err := configs.ReadConfigInto(configName, &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// logger
	logger.Open(conf.Logger.Level, conf.Logger.Stdout,
		conf.Logger.File.Name, conf.Logger.File.MaxSize, conf.Logger.File.MaxBackup,
		conf.Logger.File.MaxAge, conf.Logger.File.Compressed)
	defer logger.Close()

	// db
	sqlDB, err := si.OpenSqlDB(conf.Server.Repo.Driver, conf.Server.Repo.ConnStr,
		conf.Server.Repo.MaxIdleConns, conf.Server.Repo.MaxOpenConns,
		time.Duration(conf.Server.Repo.ConnMaxLifetimeMinutes)*time.Minute)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	// gorm
	var gormDB *gorm.DB
	switch conf.Server.Repo.Driver {
	case "pgx":
		gormDB, err = sigorm.OpenPostgres(sqlDB)
	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	var txBeginner common.TxBeginner
	var isolationLvl common.IsolationLevelSetter
	var cartRepo port.CartRepo
	var cartProductRepo port.CartProductRepo
	switch conf.Server.Repo.Driver {
	case "pgx":
		txBeginner = txcom.NewGormTxBeginner(gormDB)
		isolationLvl = txcom.NewGormIsolationLevelSetter()
		cartRepo = adapter.NewCartPg(gormDB)
		cartProductRepo = adapter.NewCartProductPg(gormDB)
	default:
		logger.Error(conf.Server.Repo.Driver + " is not allowed")
		os.Exit(1)
	}
	usc := usecase.NewCartUsc(txBeginner, isolationLvl, cartRepo, cartProductRepo)

	// oauth2
	idTokenValidators := make(commonport.IDTokenValidators)
	for _, v := range conf.Client.Oauth2.IDTokenValidators {
		if v.Type == "jwks" {
			jwksUrl, err := commonadapter.GetJwksUrl(v.OpenIDConfUrl)
			if err != nil {
				logger.Error(err.Error())
				os.Exit(1)
			}
			validator := commonadapter.NewJwksIDTokenValidator(jwksUrl, v.Token.TokenSourceKeyName, v.Token.IDKeyName, v.Token.IDTokenKeyName)
			idTokenValidators[v.Token.Source] = validator
		}
	}

	var userSvc commonport.UserSvc
	if conf.Client.UserHttp.Url != "" {
		userSvc = commonadapter.NewUserHttp(sihttp.DefaultInsecureClient(),
			// conf.Client.Oauth2.Token.Source,
			conf.Client.UserHttp.Url,
			conf.Client.UserHttp.BearerToken,
			conf.Client.Oauth2.Token.TokenSourceKeyName,
			conf.Client.Oauth2.Token.IDKeyName, conf.Client.Oauth2.Token.IDTokenKeyName)
	} else if conf.Client.UserGrpc.Addr != "" {
		conn, err := wrapper.NewGrpcClient(conf.Client.UserGrpc, false)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		userSvc = commonadapter.NewUserGrpc(conn)
	} else {
		userSvc = commonadapter.NewUserSvcNop()
	}

	// http handler
	// userHandler = delivery.NewUserHttpHandler(defaultTimeout, userUsc)
	router := mux.NewRouter()
	// SetRoute(router, conf.Server.Http, idTokenValidators)
	route.CartRoute(router, conf.Server.Http, idTokenValidators, usc, userSvc)

	// http server
	tlsConfig := sihttp.CreateTLSConfigMinTls(tls.VersionTLS12)
	httpServer := sihttp.NewServerCors(router, tlsConfig, addr,
		time.Duration(writeTimeout)*time.Second, time.Duration(readTimeout)*time.Second,
		certPem, certKey,
		strings.Split(conf.Server.Http.AllowedOrigins, ","),
		strings.Split(conf.Server.Http.AllowedHeaders, ","),
		strings.Split(conf.Server.Http.AllowedMethods, ","),
	)

	// ticker
	ticker := time.NewTicker(time.Duration(tickIntervalSec) * time.Second)
	tickerDone := make(chan bool)
	common.StartTicker(tickerDone, ticker, func(t time.Time) {
		logger.Info(fmt.Sprintf("NoOfGR:%v, %v", runtime.NumGoroutine(), t))
	})

	// signal, wait for it to shutdown http server.
	common.StartSignalStopper(httpServer, syscall.SIGINT, syscall.SIGTERM)

	// start
	logger.Info("start listening on " + addr)
	if err = httpServer.Start(); err != nil {
		logger.Error(err.Error())
	}

	// finish
	ticker.Stop()
	tickerDone <- true
	logger.Info("finished")
}

// var (
// 	userHandler *delivery.UserHttpHandler
// )

// func SetRoute(router *mux.Router, conf common.ConfigHttp, validator commonport.IDTokenValidators) {
// 	// router.HandleFunc("/v1/user/{login_source}",
// 	// 	middlewares.AuthBearerHandler(userHandler.HandleRegisterUser, conf.BearerToken),
// 	// ).Methods(http.MethodPost)

// 	router.HandleFunc("/v1/user/account",
// 		middlewares.AuthIDTokenHandler(userHandler.HandleFindByLoginID, validator, "tid", "id_token", "token_source"),
// 	).Methods(http.MethodGet)

// 	router.HandleFunc("/v1/user",
// 		middlewares.AuthBearerHandler(userHandler.HandleRegisterUser, conf.BearerToken),
// 	).Methods(http.MethodPost)

// 	router.HandleFunc("/v1/user/{id}",
// 		middlewares.AuthIDTokenHandler(userHandler.HandleFindUser, validator, "tid", "id_token", "token_source"),
// 	).Methods(http.MethodGet)

// 	router.HandleFunc("/v1/user/{id}",
// 		middlewares.AuthJWTHandler(userHandler.HandleRemoveUser, conf.Jwt.Secret),
// 	).Methods(http.MethodDelete)

// }
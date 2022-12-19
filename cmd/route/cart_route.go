package route

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/middlewares"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/order/delivery"
	"github.com/w-woong/order/port"
)

func CartRoute(router *mux.Router, conf common.ConfigHttp,
	validator commonport.IDTokenValidators, usc port.CartUsc, userSvc commonport.UserSvc) *delivery.CartHttpHandler {

	handler := delivery.NewCartHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	router.HandleFunc("/v1/order/cart", middlewares.AuthIDTokenUserAccountHandler(
		handler.HandleFindByUserID, validator, userSvc,
	)).Methods(http.MethodGet)
	router.HandleFunc("/v1/order/cart/_find-or-create", middlewares.AuthIDTokenUserAccountHandler(
		handler.HandleFindOrCreateByUserID, validator, userSvc,
	)).Methods(http.MethodGet)

	router.HandleFunc("/test/order/cart", middlewares.AuthIDTokenUserAccountHandler(
		handler.HandleTestRefreshError, validator, userSvc,
	)).Methods(http.MethodGet)

	router.HandleFunc("/v1/order/cart/product", middlewares.AuthIDTokenUserAccountHandler(
		handler.HandleAddCartProduct, validator, userSvc,
	)).Methods(http.MethodPost)

	return handler
}

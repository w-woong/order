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
	tokenCookie commonport.TokenCookie,
	parser commonport.IDTokenParser, usc port.CartUsc, userSvc commonport.UserSvc) *delivery.CartHttpHandler {

	handler := delivery.NewCartHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	router.HandleFunc("/v1/order/cart", middlewares.AuthIDTokenUserAccountSvc(
		handler.HandleFindByUserID, tokenCookie, parser, userSvc,
	)).Methods(http.MethodGet)
	router.HandleFunc("/v1/order/cart/_find-or-create", middlewares.AuthIDTokenUserAccountSvc(
		handler.HandleFindOrCreateByUserID, tokenCookie, parser, userSvc,
	)).Methods(http.MethodGet)

	router.HandleFunc("/test/order/cart", middlewares.AuthIDTokenUserAccountSvc(
		handler.HandleTestRefreshError, tokenCookie, parser, userSvc,
	)).Methods(http.MethodGet)

	router.HandleFunc("/v1/order/cart/product", middlewares.AuthIDTokenUserAccountSvc(
		handler.HandleAddCartProduct, tokenCookie, parser, userSvc,
	)).Methods(http.MethodPost)

	return handler
}

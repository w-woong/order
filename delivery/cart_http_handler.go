package delivery

import (
	"bytes"
	"net/http"
	"time"

	"github.com/go-wonk/si"
	"github.com/w-woong/common"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/order/dto"
	"github.com/w-woong/order/port"
)

type CartHttpHandler struct {
	timeout time.Duration
	usc     port.CartUsc
}

func NewCartHttpHandler(timeout time.Duration, usc port.CartUsc) *CartHttpHandler {
	return &CartHttpHandler{
		timeout: timeout,
		usc:     usc,
	}
}

func (d *CartHttpHandler) HandleFindByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// claims, ok := ctx.Value(commondto.IDTokenClaimsKey{}).(commondto.IDTokenClaims)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	// tokenSource, ok := ctx.Value(commondto.TokenSourceKey{}).(string)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	user, ok := ctx.Value(commondto.UserAccountKey{}).(commondto.User)
	if !ok {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("could not find claims", logger.UrlField(r.URL.String()))
		return
	}

	cart, err := d.usc.FindByUserID(ctx, user.ID)
	if err != nil {
		common.HttpErrorWithBody(w, http.StatusInternalServerError, common.NewHttpBody(err.Error(), http.StatusInternalServerError))
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
		Document: &cart,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

}

func (d *CartHttpHandler) HandleFindOrCreateByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// claims, ok := ctx.Value(commondto.IDTokenClaimsKey{}).(commondto.IDTokenClaims)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	// tokenSource, ok := ctx.Value(commondto.TokenSourceKey{}).(string)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	user, ok := ctx.Value(commondto.UserAccountKey{}).(commondto.User)
	if !ok {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("could not find claims", logger.UrlField(r.URL.String()))
		return
	}

	cart, err := d.usc.FindOrCreateByUserID(ctx, user.ID)
	if err != nil {
		common.HttpErrorWithBody(w, http.StatusInternalServerError, common.NewHttpBody(err.Error(), http.StatusInternalServerError))
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
		Document: &cart,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

}

func (d *CartHttpHandler) HandleAddCartProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// claims, ok := ctx.Value(commondto.IDTokenClaimsKey{}).(commondto.IDTokenClaims)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	// tokenSource, ok := ctx.Value(commondto.TokenSourceKey{}).(string)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	user, ok := ctx.Value(commondto.UserAccountKey{}).(commondto.User)
	if !ok {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("could not find user", logger.UrlField(r.URL.String()))
		return
	}

	cartProduct := dto.CartProduct{}
	reqBody := common.HttpBody{
		Document: &cartProduct,
	}
	var copiedReqBody *bytes.Buffer
	var err error
	if copiedReqBody, err = si.DecodeJsonCopied(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()), logger.ReqBodyField(copiedReqBody.Bytes()))
		return
	}

	cart, err := d.usc.FindByUserID(ctx, user.ID)
	if err != nil {
		common.HttpErrorWithBody(w, http.StatusInternalServerError, common.NewHttpBody(err.Error(), http.StatusInternalServerError))
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	rowsAffected, err := d.usc.AddCartProduct(ctx, cart.ID, cartProduct)
	if err != nil {
		common.HttpErrorWithBody(w, http.StatusInternalServerError, common.NewHttpBody(err.Error(), http.StatusInternalServerError))
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status: http.StatusOK,
		Count:  int(rowsAffected),
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

}

func (d *CartHttpHandler) HandleFindByUserIDError(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// claims, ok := ctx.Value(commondto.IDTokenClaimsKey{}).(commondto.IDTokenClaims)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	// tokenSource, ok := ctx.Value(commondto.TokenSourceKey{}).(string)
	// if !ok {
	// 	common.HttpError(w, http.StatusInternalServerError)
	// 	logger.Error("could not find claims", logger.UrlField(r.URL.String()))
	// 	return
	// }

	user, ok := ctx.Value(commondto.UserAccountKey{}).(commondto.User)
	if !ok {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("could not find claims", logger.UrlField(r.URL.String()))
		return
	}

	cart, err := d.usc.FindByUserID(ctx, user.ID)
	if err != nil {
		common.HttpErrorWithBody(w, http.StatusInternalServerError, common.NewHttpBody(err.Error(), http.StatusInternalServerError))
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   common.StatusTryRefreshIDToken,
		Count:    1,
		Document: &cart,
	}

	common.HttpErrorWithBody(w, http.StatusUnauthorized, &resBody)
	logger.Error("error test", logger.UrlField(r.URL.String()))

}

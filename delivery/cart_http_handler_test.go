package delivery_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/order/delivery"
	"github.com/w-woong/order/dto"
	"github.com/w-woong/order/port/mocks"
)

var (
	userDto = commondto.User{
		ID: "user_id_1234",
	}
	cartID = "cart_id_1234"
	cart   = dto.Cart{
		ID: cartID,
	}
	cartProduct = dto.CartProduct{
		ProductID: "product_id_1234",
		Cost:      1000,
		Qty:       1,
		Amt:       1000,
		Selected:  true,
	}
)

func Test_CartHttpHandler_HandleAddCartProduct(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, commondto.UserAccountContextKey{}, userDto)

	urlPath := "/v1/order/cart/product"
	ctrl := gomock.NewController(t)

	usc := mocks.NewMockCartUsc(ctrl)
	usc.EXPECT().FindByUserID(gomock.Any(), userDto.ID).Return(cart, nil).AnyTimes()
	usc.EXPECT().AddCartProduct(gomock.Any(), cartID, cartProduct).Return(int64(1), nil).AnyTimes()

	router := mux.NewRouter()
	handler := delivery.NewCartHttpHandler(15*time.Second, usc)
	router.HandleFunc(urlPath, handler.HandleAddCartProduct).Methods(http.MethodPost)

	reqBody := common.HttpBody{
		Document: &cartProduct,
	}
	b, _ := json.Marshal(&reqBody)
	body := bytes.NewBuffer(b)
	fmt.Println(reqBody.String())

	req, err := http.NewRequest(http.MethodPost, urlPath, body)
	assert.Nil(t, err)
	req = req.WithContext(ctx)

	// response 받을 레코더 초기화
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	resBody, err := io.ReadAll(rr.Body)
	assert.Nil(t, err)
	fmt.Println(string(resBody))
	assert.Equal(t, http.StatusOK, rr.Code)
}

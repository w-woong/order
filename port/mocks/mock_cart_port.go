// Code generated by MockGen. DO NOT EDIT.
// Source: ./cart_port.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/w-woong/common"
	dto "github.com/w-woong/order/dto"
	entity "github.com/w-woong/order/entity"
)

// MockCartRepo is a mock of CartRepo interface.
type MockCartRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepoMockRecorder
}

// MockCartRepoMockRecorder is the mock recorder for MockCartRepo.
type MockCartRepoMockRecorder struct {
	mock *MockCartRepo
}

// NewMockCartRepo creates a new mock instance.
func NewMockCartRepo(ctrl *gomock.Controller) *MockCartRepo {
	mock := &MockCartRepo{ctrl: ctrl}
	mock.recorder = &MockCartRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepo) EXPECT() *MockCartRepoMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockCartRepo) CreateCart(ctx context.Context, tx common.TxController, cart entity.Cart) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", ctx, tx, cart)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockCartRepoMockRecorder) CreateCart(ctx, tx, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockCartRepo)(nil).CreateCart), ctx, tx, cart)
}

// DeleteCart mocks base method.
func (m *MockCartRepo) DeleteCart(ctx context.Context, tx common.TxController, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", ctx, tx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockCartRepoMockRecorder) DeleteCart(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockCartRepo)(nil).DeleteCart), ctx, tx, id)
}

// ReadByUserID mocks base method.
func (m *MockCartRepo) ReadByUserID(ctx context.Context, tx common.TxController, userID string) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByUserID", ctx, tx, userID)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByUserID indicates an expected call of ReadByUserID.
func (mr *MockCartRepoMockRecorder) ReadByUserID(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByUserID", reflect.TypeOf((*MockCartRepo)(nil).ReadByUserID), ctx, tx, userID)
}

// ReadByUserIDNoTx mocks base method.
func (m *MockCartRepo) ReadByUserIDNoTx(ctx context.Context, userID string) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByUserIDNoTx", ctx, userID)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByUserIDNoTx indicates an expected call of ReadByUserIDNoTx.
func (mr *MockCartRepoMockRecorder) ReadByUserIDNoTx(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByUserIDNoTx", reflect.TypeOf((*MockCartRepo)(nil).ReadByUserIDNoTx), ctx, userID)
}

// ReadCart mocks base method.
func (m *MockCartRepo) ReadCart(ctx context.Context, tx common.TxController, id string) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCart", ctx, tx, id)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCart indicates an expected call of ReadCart.
func (mr *MockCartRepoMockRecorder) ReadCart(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCart", reflect.TypeOf((*MockCartRepo)(nil).ReadCart), ctx, tx, id)
}

// ReadCartNoTx mocks base method.
func (m *MockCartRepo) ReadCartNoTx(ctx context.Context, id string) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCartNoTx", ctx, id)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCartNoTx indicates an expected call of ReadCartNoTx.
func (mr *MockCartRepoMockRecorder) ReadCartNoTx(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCartNoTx", reflect.TypeOf((*MockCartRepo)(nil).ReadCartNoTx), ctx, id)
}

// UpdateCart mocks base method.
func (m *MockCartRepo) UpdateCart(ctx context.Context, tx common.TxController, id string, cart entity.Cart) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, tx, id, cart)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockCartRepoMockRecorder) UpdateCart(ctx, tx, id, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockCartRepo)(nil).UpdateCart), ctx, tx, id, cart)
}

// MockCartProductRepo is a mock of CartProductRepo interface.
type MockCartProductRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCartProductRepoMockRecorder
}

// MockCartProductRepoMockRecorder is the mock recorder for MockCartProductRepo.
type MockCartProductRepoMockRecorder struct {
	mock *MockCartProductRepo
}

// NewMockCartProductRepo creates a new mock instance.
func NewMockCartProductRepo(ctrl *gomock.Controller) *MockCartProductRepo {
	mock := &MockCartProductRepo{ctrl: ctrl}
	mock.recorder = &MockCartProductRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartProductRepo) EXPECT() *MockCartProductRepoMockRecorder {
	return m.recorder
}

// CreateCartProduct mocks base method.
func (m *MockCartProductRepo) CreateCartProduct(ctx context.Context, tx common.TxController, cartProduct entity.CartProduct) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartProduct", ctx, tx, cartProduct)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCartProduct indicates an expected call of CreateCartProduct.
func (mr *MockCartProductRepoMockRecorder) CreateCartProduct(ctx, tx, cartProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartProduct", reflect.TypeOf((*MockCartProductRepo)(nil).CreateCartProduct), ctx, tx, cartProduct)
}

// DeleteByCartID mocks base method.
func (m *MockCartProductRepo) DeleteByCartID(ctx context.Context, tx common.TxController, cartID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByCartID", ctx, tx, cartID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByCartID indicates an expected call of DeleteByCartID.
func (mr *MockCartProductRepoMockRecorder) DeleteByCartID(ctx, tx, cartID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByCartID", reflect.TypeOf((*MockCartProductRepo)(nil).DeleteByCartID), ctx, tx, cartID)
}

// DeleteCartProduct mocks base method.
func (m *MockCartProductRepo) DeleteCartProduct(ctx context.Context, tx common.TxController, id string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartProduct", ctx, tx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCartProduct indicates an expected call of DeleteCartProduct.
func (mr *MockCartProductRepoMockRecorder) DeleteCartProduct(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartProduct", reflect.TypeOf((*MockCartProductRepo)(nil).DeleteCartProduct), ctx, tx, id)
}

// ReadCartProduct mocks base method.
func (m *MockCartProductRepo) ReadCartProduct(ctx context.Context, tx common.TxController, id string) (entity.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCartProduct", ctx, tx, id)
	ret0, _ := ret[0].(entity.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCartProduct indicates an expected call of ReadCartProduct.
func (mr *MockCartProductRepoMockRecorder) ReadCartProduct(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCartProduct", reflect.TypeOf((*MockCartProductRepo)(nil).ReadCartProduct), ctx, tx, id)
}

// ReadCartProductNoTx mocks base method.
func (m *MockCartProductRepo) ReadCartProductNoTx(ctx context.Context, id string) (entity.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCartProductNoTx", ctx, id)
	ret0, _ := ret[0].(entity.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCartProductNoTx indicates an expected call of ReadCartProductNoTx.
func (mr *MockCartProductRepoMockRecorder) ReadCartProductNoTx(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCartProductNoTx", reflect.TypeOf((*MockCartProductRepo)(nil).ReadCartProductNoTx), ctx, id)
}

// UpdateCartProduct mocks base method.
func (m *MockCartProductRepo) UpdateCartProduct(ctx context.Context, tx common.TxController, id string, cartProduct entity.CartProduct) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartProduct", ctx, tx, id, cartProduct)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartProduct indicates an expected call of UpdateCartProduct.
func (mr *MockCartProductRepoMockRecorder) UpdateCartProduct(ctx, tx, id, cartProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartProduct", reflect.TypeOf((*MockCartProductRepo)(nil).UpdateCartProduct), ctx, tx, id, cartProduct)
}

// MockCartUsc is a mock of CartUsc interface.
type MockCartUsc struct {
	ctrl     *gomock.Controller
	recorder *MockCartUscMockRecorder
}

// MockCartUscMockRecorder is the mock recorder for MockCartUsc.
type MockCartUscMockRecorder struct {
	mock *MockCartUsc
}

// NewMockCartUsc creates a new mock instance.
func NewMockCartUsc(ctrl *gomock.Controller) *MockCartUsc {
	mock := &MockCartUsc{ctrl: ctrl}
	mock.recorder = &MockCartUscMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartUsc) EXPECT() *MockCartUscMockRecorder {
	return m.recorder
}

// AddCartProduct mocks base method.
func (m *MockCartUsc) AddCartProduct(ctx context.Context, cartID string, cartProduct dto.CartProduct) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCartProduct", ctx, cartID, cartProduct)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCartProduct indicates an expected call of AddCartProduct.
func (mr *MockCartUscMockRecorder) AddCartProduct(ctx, cartID, cartProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCartProduct", reflect.TypeOf((*MockCartUsc)(nil).AddCartProduct), ctx, cartID, cartProduct)
}

// FindByUserID mocks base method.
func (m *MockCartUsc) FindByUserID(ctx context.Context, userID string) (dto.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].(dto.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *MockCartUscMockRecorder) FindByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockCartUsc)(nil).FindByUserID), ctx, userID)
}

// FindOrCreateByUserID mocks base method.
func (m *MockCartUsc) FindOrCreateByUserID(ctx context.Context, userID string) (dto.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreateByUserID", ctx, userID)
	ret0, _ := ret[0].(dto.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreateByUserID indicates an expected call of FindOrCreateByUserID.
func (mr *MockCartUscMockRecorder) FindOrCreateByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreateByUserID", reflect.TypeOf((*MockCartUsc)(nil).FindOrCreateByUserID), ctx, userID)
}

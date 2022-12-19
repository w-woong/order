package port

//go:generate mockgen -destination=./mocks/mock_cart_port.go -package=mocks -mock_names=CartRepo=MockCartRepo,CartProductRepo=MockCartProductRepo,CartUsc=MockCartUsc -source=./cart_port.go . CartRepo,CartProductRepo,CartUsc

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/order/dto"
	"github.com/w-woong/order/entity"
)

type CartRepo interface {
	CreateCart(ctx context.Context, tx common.TxController, cart entity.Cart) (int64, error)
	ReadCart(ctx context.Context, tx common.TxController, id string) (entity.Cart, error)
	ReadCartNoTx(ctx context.Context, id string) (entity.Cart, error)
	UpdateCart(ctx context.Context, tx common.TxController, id string, cart entity.Cart) (int64, error)
	DeleteCart(ctx context.Context, tx common.TxController, id string) (int64, error)

	ReadByUserID(ctx context.Context, tx common.TxController, userID string) (entity.Cart, error)
	ReadByUserIDNoTx(ctx context.Context, userID string) (entity.Cart, error)
}

type CartProductRepo interface {
	CreateCartProduct(ctx context.Context, tx common.TxController,
		cartProduct entity.CartProduct) (int64, error)
	ReadCartProduct(ctx context.Context, tx common.TxController, id string) (entity.CartProduct, error)
	ReadCartProductNoTx(ctx context.Context, id string) (entity.CartProduct, error)
	UpdateCartProduct(ctx context.Context, tx common.TxController,
		id string, cartProduct entity.CartProduct) (int64, error)
	DeleteCartProduct(ctx context.Context, tx common.TxController, id string) (int64, error)
	DeleteByCartID(ctx context.Context, tx common.TxController, cartID string) (int64, error)
}

type CartUsc interface {
	FindByUserID(ctx context.Context, userID string) (dto.Cart, error)
	FindOrCreateByUserID(ctx context.Context, userID string) (dto.Cart, error)
	AddCartProduct(ctx context.Context, cartID string, cartProduct dto.CartProduct) (int64, error)
	// ModifyCartProduct(ctx context.Context, id string, cartProduct dto.CartProduct) (int64, error)
	// RemoveCartProduct(ctx context.Context, id string) (int64, error)
	// ClearCart(ctx context.Context, id string) (int64, error)
}

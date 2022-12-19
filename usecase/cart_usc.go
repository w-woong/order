package usecase

import (
	"context"
	"errors"

	"github.com/w-woong/common"
	"github.com/w-woong/order/conv"
	"github.com/w-woong/order/dto"
	"github.com/w-woong/order/entity"
	"github.com/w-woong/order/port"
)

type cartUsc struct {
	beginner        common.TxBeginner
	isolationLvl    common.IsolationLevelSetter
	cartRepo        port.CartRepo
	cartProductRepo port.CartProductRepo
}

func NewCartUsc(beginner common.TxBeginner, isolationLvl common.IsolationLevelSetter, cartRepo port.CartRepo, cartProductRepo port.CartProductRepo) *cartUsc {
	return &cartUsc{
		beginner:        beginner,
		isolationLvl:    isolationLvl,
		cartRepo:        cartRepo,
		cartProductRepo: cartProductRepo,
	}
}

func (u *cartUsc) AddCartProduct(ctx context.Context, cartID string, cartProduct dto.CartProduct) (int64, error) {
	e, err := conv.ToCartProductEntity(&cartProduct)
	if err != nil {
		return 0, err
	}

	e.CreateSetID()
	e.ReferTo(cartID)

	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rowsAffected, err := u.cartProductRepo.CreateCartProduct(ctx, tx, e)
	if err != nil {
		return 0, err
	}

	return rowsAffected, tx.Commit()
}

func (u *cartUsc) FindByUserID(ctx context.Context, userID string) (dto.Cart, error) {
	cart, err := u.cartRepo.ReadByUserIDNoTx(ctx, userID)
	if err != nil {
		return dto.NilCart, err
	}

	return conv.ToCartDto(&cart)
}

func (u *cartUsc) FindOrCreateByUserID(ctx context.Context, userID string) (dto.Cart, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return dto.NilCart, err
	}
	defer tx.Rollback()

	if err = u.isolationLvl.SetSerializable(tx); err != nil {
		return dto.NilCart, err
	}

	cart, err := u.cartRepo.ReadByUserID(ctx, tx, userID)
	if err != nil {
		if !errors.Is(err, common.ErrRecordNotFound) {
			return dto.NilCart, err
		}
	}

	if errors.Is(err, common.ErrRecordNotFound) {
		_, err = u.makeCartWithTx(ctx, tx, userID)
		if err != nil {
			return dto.NilCart, err
		}
		cart, err = u.cartRepo.ReadByUserID(ctx, tx, userID)
		if err != nil {
			return dto.NilCart, err
		}
	}

	cartDto, err := conv.ToCartDto(&cart)
	if err != nil {
		return dto.NilCart, err
	}
	return cartDto, tx.Commit()
}
func (u *cartUsc) makeCart(ctx context.Context, userID string) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rowsAffected, err := u.makeCartWithTx(ctx, tx, userID)
	if err != nil {
		return 0, err
	}

	return rowsAffected, tx.Commit()

}

func (u *cartUsc) makeCartWithTx(ctx context.Context, tx common.TxController, userID string) (int64, error) {
	cart := entity.Cart{}
	cart.CreateSetID()
	cart.ReferTo(userID)

	rowsAffected, err := u.cartRepo.CreateCart(ctx, tx, cart)

	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New("failed to create a new cart")
	}

	return rowsAffected, nil

}

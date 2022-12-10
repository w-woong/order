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
	cartRepo        port.CartRepo
	cartProductRepo port.CartProductRepo
}

func NewCartUsc(beginner common.TxBeginner, cartRepo port.CartRepo, cartProductRepo port.CartProductRepo) *cartUsc {
	return &cartUsc{
		beginner:        beginner,
		cartRepo:        cartRepo,
		cartProductRepo: cartProductRepo,
	}
}

func (u *cartUsc) FindByUserID(ctx context.Context, userID string) (dto.Cart, error) {
	cart, err := u.cartRepo.ReadByUserIDNoTx(ctx, userID)
	if err != nil {
		if !errors.Is(err, common.ErrRecordNotFound) {
			return dto.NilCart, err
		}
	}

	if errors.Is(err, common.ErrRecordNotFound) {
		_, err = u.makeCart(ctx, userID)
		if err != nil {
			return dto.NilCart, err
		}
		cart, err = u.cartRepo.ReadByUserIDNoTx(ctx, userID)
		if err != nil {
			return dto.NilCart, err
		}
	}

	return conv.ToCartDto(&cart)
}

func (u *cartUsc) makeCart(ctx context.Context, userID string) (int64, error) {
	tx, err := u.beginner.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	rowsAffected, err := u.cartRepo.CreateCart(ctx, tx, entity.Cart{
		UserID: userID,
	})

	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New("failed to create a new cart")
	}

	return rowsAffected, tx.Commit()

}

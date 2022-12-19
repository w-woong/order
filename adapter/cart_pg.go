package adapter

import (
	"context"

	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/order/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cartPg struct {
	db *gorm.DB
}

func NewCartPg(db *gorm.DB) *cartPg {
	return &cartPg{
		db: db,
	}
}

func (a *cartPg) CreateCart(ctx context.Context, tx common.TxController, cart entity.Cart) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&cart)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *cartPg) ReadCart(ctx context.Context, tx common.TxController, id string) (entity.Cart, error) {

	return a.readCart(ctx, tx.(*txcom.GormTxController).Tx, id)
}

func (a *cartPg) ReadCartNoTx(ctx context.Context, id string) (entity.Cart, error) {
	return a.readCart(ctx, a.db, id)
}

func (a *cartPg) readCart(ctx context.Context, db *gorm.DB, id string) (entity.Cart, error) {

	var cart entity.Cart
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Preload(clause.Associations).
		Limit(1).Find(&cart)
	if res.Error != nil {
		return entity.NilCart, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilCart, common.ErrRecordNotFound
	}
	return cart, nil
}

func (a *cartPg) ReadByUserID(ctx context.Context, tx common.TxController, userID string) (entity.Cart, error) {
	return a.readByUserID(ctx, tx.(*txcom.GormTxController).Tx, userID)
}
func (a *cartPg) ReadByUserIDNoTx(ctx context.Context, userID string) (entity.Cart, error) {
	return a.readByUserID(ctx, a.db, userID)
}

func (a *cartPg) UpdateCart(ctx context.Context, tx common.TxController, id string, cart entity.Cart) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		// Session(&gorm.Session{FullSaveAssociations: true}). // update all assossications
		WithContext(ctx).
		Model(&cart).
		Where("id = ?", id).
		Updates(&cart)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *cartPg) DeleteCart(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Delete(&entity.Cart{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *cartPg) readByUserID(ctx context.Context, db *gorm.DB, userID string) (entity.Cart, error) {

	var cart entity.Cart
	res := db.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload(clause.Associations).
		Limit(1).Find(&cart)
	if res.Error != nil {
		return entity.NilCart, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilCart, common.ErrRecordNotFound
	}
	return cart, nil
}

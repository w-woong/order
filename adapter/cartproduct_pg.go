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

type cartProductPg struct {
	db *gorm.DB
}

func NewCartProductPg(db *gorm.DB) *cartProductPg {
	return &cartProductPg{
		db: db,
	}
}

func (a *cartProductPg) CreateCartProduct(ctx context.Context, tx common.TxController, cartProduct entity.CartProduct) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.WithContext(ctx).
		Create(&cartProduct)

	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *cartProductPg) ReadCartProduct(ctx context.Context, tx common.TxController, id string) (entity.CartProduct, error) {

	return a.readCartProduct(ctx, tx.(*txcom.GormTxController).Tx, id)
}

func (a *cartProductPg) ReadCartProductNoTx(ctx context.Context, id string) (entity.CartProduct, error) {
	return a.readCartProduct(ctx, a.db, id)
}

func (a *cartProductPg) readCartProduct(ctx context.Context, db *gorm.DB, id string) (entity.CartProduct, error) {

	var cartProduct entity.CartProduct
	res := db.WithContext(ctx).
		Where("id = ?", id).
		Preload(clause.Associations).
		Limit(1).Find(&cartProduct)
	if res.Error != nil {
		return entity.NilCartProduct, txcom.ConvertErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.NilCartProduct, common.ErrRecordNotFound
	}
	return cartProduct, nil
}

func (a *cartProductPg) UpdateCartProduct(ctx context.Context, tx common.TxController, id string, cartProduct entity.CartProduct) (int64, error) {

	res := tx.(*txcom.GormTxController).Tx.
		// Session(&gorm.Session{FullSaveAssociations: true}). // update all assossications
		WithContext(ctx).
		Model(&cartProduct).
		Where("id = ?", id).
		Updates(&cartProduct)
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}

	return res.RowsAffected, nil
}
func (a *cartProductPg) DeleteCartProduct(ctx context.Context, tx common.TxController, id string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Delete(&entity.CartProduct{ID: id})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

func (a *cartProductPg) DeleteByCartID(ctx context.Context, tx common.TxController, cartID string) (int64, error) {
	res := tx.(*txcom.GormTxController).Tx.
		WithContext(ctx).
		Where("cart_id = ?", cartID).
		Delete(&entity.CartProduct{})
	if res.Error != nil {
		logger.Error(res.Error.Error())
		return 0, txcom.ConvertErr(res.Error)
	}
	return res.RowsAffected, nil
}

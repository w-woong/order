package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common/txcom"
	"github.com/w-woong/order/adapter"
	"github.com/w-woong/order/entity"
)

var (
	cart0 = entity.Cart{
		ID:     "cart0000",
		UserID: "user0000",
		CartProducts: entity.CartProductList{
			entity.CartProduct{
				ID:        "cartproduct0",
				CartID:    "cart0000",
				ProductID: "product0",
			},
		},
	}
)

func Test_cartPg_CreateProduct(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	beginner := txcom.NewGormTxBeginner(gdb)
	repo := adapter.NewCartPg(gdb)

	tx, err := beginner.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	affected, err := repo.CreateCart(ctx, tx, cart0)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, affected)

	assert.Nil(t, tx.Commit())
}

func Test_cartPg_ReadProductNoTx(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	repo := adapter.NewCartPg(gdb)

	res, err := repo.ReadCartNoTx(ctx, cart0.ID)
	assert.Nil(t, err)
	// assert.EqualValues(t, 1, affected)
	fmt.Println(res.String())

}

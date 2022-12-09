package dto

import (
	"encoding/json"
	"time"
)

var (
	NilCart        = Cart{}
	NilCartProduct = CartProduct{}
)

type Cart struct {
	ID           string          `json:"id"`
	CreatedAt    *time.Time      `json:"created_at,omitempty"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty"`
	UserID       string          `json:"user_id"`
	CartProducts CartProductList `json:"card_products"`
}

func (e *Cart) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type CartList []Cart

func (e *CartList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type CartProduct struct {
	ID         string     `json:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	CartID     string     `json:"cart_id"`
	ProductID  string     `json:"product_id"`
	ProductQty int        `json:"product_qty"`
}

func (e *CartProduct) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type CartProductList []CartProduct

func (e *CartProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

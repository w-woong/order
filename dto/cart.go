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
	ID           string          `json:"id,omitempty"`
	CreatedAt    *time.Time      `json:"created_at,omitempty"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty"`
	UserID       string          `json:"user_id,omitempty"`
	CartProducts CartProductList `json:"cart_products,omitempty"`
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
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CartID    string     `json:"cart_id,omitempty"`
	ProductID string     `json:"product_id,omitempty"`
	Cost      float64    `json:"cost,omitempty"`
	Qty       float64    `json:"qty,omitempty"`
	Amt       float64    `json:"amt,omitempty"`
	Selected  bool       `json:"selected,omitempty"`
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

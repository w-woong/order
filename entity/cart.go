package entity

import (
	"encoding/json"
	"time"
)

var (
	NilCart        = Cart{}
	NilCartProduct = CartProduct{}
)

type Cart struct {
	ID           string          `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt    *time.Time      `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty" gorm:"<-"`
	UserID       string          `json:"user_id" gorm:"column:user_id;type:string;size:64;index:idx_carts_1"`
	CartProducts CartProductList `json:"card_products" gorm:"foreignKey:CartID;references:ID"`
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
	ID         string     `json:"id" gorm:"primaryKey;type:string;size:64"`
	CreatedAt  *time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"<-"`
	CartID     string     `json:"cart_id" gorm:"not null;column:cart_id;type:string;size:64"`
	ProductID  string     `json:"product_id" gorm:"not null;column:product_id;type:string;size:64"`
	ProductQty int        `json:"product_qty" gorm:"not null;column:product_qty;type:int;scale:16;precision:4"`
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

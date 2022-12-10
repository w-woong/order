package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

var (
	NilCart        = Cart{}
	NilCartProduct = CartProduct{}
)

type Cart struct {
	ID           string          `gorm:"primaryKey;type:string;size:64" json:"id"`
	CreatedAt    *time.Time      `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt    *time.Time      `gorm:"<-" json:"updated_at,omitempty"`
	UserID       string          `gorm:"unique;column:user_id;type:string;size:64;index:idx_carts_1" json:"user_id"`
	CartProducts CartProductList `gorm:"foreignKey:CartID;references:ID" json:"card_products"`
}

func (e *Cart) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e Cart) IsNil() bool {
	return e.ID == "" && e.UserID == ""
}

func (e Cart) CreateID() string {
	return uuid.New().String()
}

func (e *Cart) CreateSetID() {
	e.ID = e.CreateID()
}

type CartList []Cart

func (e *CartList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type CartProduct struct {
	ID        string     `gorm:"primaryKey;type:string;size:64" json:"id"`
	CreatedAt *time.Time `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at,omitempty"`
	CartID    string     `gorm:"not null;column:cart_id;type:string;size:64" json:"cart_id"`
	ProductID string     `gorm:"index:idx_cart_products_1;not null;column:product_id;type:string;size:64" json:"product_id"`
	Cost      float64    `gorm:"not null;column:cost;type:float;scale:4;precision:16" json:"cost"`
	Qty       float64    `gorm:"not null;column:qty;type:float;scale:4;precision:16" json:"qty"`
	Amt       float64    `gorm:"not null;column:amt;type:float;scale:4;precision:20" json:"amt"`
	Selected  bool       `gorm:"not null;column:selected:type:bool" json:"selected"`
}

func (e *CartProduct) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
func (e CartProduct) IsNil() bool {
	return e.ID == "" && e.CartID == "" && e.ProductID == ""
}
func (e CartProduct) CreateID() string {
	return uuid.New().String()
}

func (e *CartProduct) CreateSetID() {
	e.ID = e.CreateID()
}

func (e *CartProduct) ReferTo(cartID string) {
	e.CartID = cartID
}

type CartProductList []CartProduct

func (e *CartProductList) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *CartProductList) CreateSetID() {
	for i := range *e {
		(*e)[i].CreateSetID()
	}
}

func (e *CartProductList) ReferTo(cartID string) {
	for i := range *e {
		(*e)[i].ReferTo(cartID)
	}
}

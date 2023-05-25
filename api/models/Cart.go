package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID   string    `json:"user_id"`
	Products []Product `gorm:"foreignkey:OrderID" json:"products"`
}

func (c *Cart) Prepare() {
}

func (c *Cart) Validate() error {
	if c.UserID == "" {
		return errors.New("Required UserID")
	}
	if len(c.Products) == 0 {
		return errors.New("No products in the cart")
	}
	// Puedes agregar lógica de validación adicional según tus necesidades
	return nil
}

func (c *Cart) SaveCart(db *gorm.DB) (*Cart, error) {
	var err error
	err = db.Debug().Create(&c).Error
	if err != nil {
		return &Cart{}, err
	}
	return c, nil
}

func (c *Cart) FindAllCarts(db *gorm.DB) (*[]Cart, error) {
	var err error
	carts := []Cart{}
	err = db.Debug().Model(&Cart{}).Limit(100).Find(&carts).Error
	if err != nil {
		return &[]Cart{}, err
	}
	return &carts, nil
}

func (c *Cart) FindCartByID(db *gorm.DB, cartID uint32) (*Cart, error) {
	var err error
	err = db.Debug().Model(Cart{}).Where("id = ?", cartID).Take(&c).Error
	if err != nil {
		return &Cart{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cart{}, errors.New("Cart Not Found")
	}
	return c, nil
}

func (c *Cart) UpdateCart(db *gorm.DB) (*Cart, error) {
	db = db.Debug().Model(&Cart{}).Where("id = ?", c.ID).Take(&Cart{}).UpdateColumns(
		map[string]interface{}{
			"user_id":    c.UserID,
			"products":   c.Products,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Cart{}, db.Error
	}
	err := db.Debug().Model(&Cart{}).Where("id = ?", c.ID).Take(&c).Error
	if err != nil {
		return &Cart{}, err
	}
	return c, nil
}

func (c *Cart) DeleteCart(db *gorm.DB) error {
	if err := db.Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

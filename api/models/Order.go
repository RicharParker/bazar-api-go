package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	Products  []Product `gorm:"foreignkey:OrderID" json:"products"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Status    bool      `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (o *Order) Prepare() {
	o.ID = 0
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}

func (o *Order) Validate() error {
	if o.UserID == "" {
		return errors.New("Required UserID")
	}
	if len(o.Products) == 0 {
		return errors.New("No products in the order")
	}
	if o.Amount == 0 {
		return errors.New("Required Amount")
	}
	// Additional validation logic for the order
	return nil
}

func (o *Order) SaveOrder(db *gorm.DB) (*Order, error) {
	var err error
	err = db.Debug().Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}

func (o *Order) FindAllOrders(db *gorm.DB) (*[]Order, error) {
	var err error
	orders := []Order{}
	err = db.Debug().Model(&Order{}).Limit(100).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	return &orders, nil
}

func (o *Order) FindOrderByID(db *gorm.DB, oid uint32) (*Order, error) {
	var err error
	err = db.Debug().Model(Order{}).Where("id = ?", oid).Take(&o).Error
	if err != nil {
		return &Order{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Order{}, errors.New("Order Not Found")
	}
	return o, nil
}

func (o *Order) UpdateOrder(db *gorm.DB) (*Order, error) {
	db = db.Debug().Model(&Order{}).Where("id = ?", o.ID).Take(&Order{}).UpdateColumns(
		map[string]interface{}{
			"user_id":    o.UserID,
			"products":   o.Products,
			"amount":     o.Amount,
			"status":     o.Status,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Order{}, db.Error
	}
	err := db.Debug().Model(&Order{}).Where("id = ?", o.ID).Take(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}

func (o *Order) DeleteOrder(db *gorm.DB) error {
	if err := db.Delete(&o).Error; err != nil {
		return err
	}
	return nil
}

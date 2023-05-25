package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	Desc       string    `gorm:"text" json:"desc"`
	Img        string    `gorm:"size:255" json:"img"`
	Categories []string  `gorm:"-" json:"categories"`
	Size       []string  `gorm:"-" json:"size"`
	Color      []string  `gorm:"-" json:"color"`
	Price      float64   `gorm:"not null" json:"price"`
	InStock    bool      `gorm:"not null" json:"in_stock"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Product) Prepare() {
	u.ID = 0
	u.Title = html.EscapeString(strings.TrimSpace(u.Title))
	u.Desc = html.EscapeString(strings.TrimSpace(u.Desc))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Product) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Title == "" {
			return errors.New("Required Title")
		}
		if u.Price == 0 {
			return errors.New("Required Price")
		}
		if err := validateCategories(u.Categories); err != nil {
			return err
		}
		if err := validateSize(u.Size); err != nil {
			return err
		}
		if err := validateColor(u.Color); err != nil {
			return err
		}
		return nil
	case "login":
		// No validations required for login action
		return nil
	default:
		if u.Title == "" {
			return errors.New("Required Title")
		}
		if u.Price == 0 {
			return errors.New("Required Price")
		}
		if err := validateCategories(u.Categories); err != nil {
			return err
		}
		if err := validateSize(u.Size); err != nil {
			return err
		}
		if err := validateColor(u.Color); err != nil {
			return err
		}
		return nil
	}
}

func validateCategories(categories []string) error {

	return nil
}

func validateSize(size []string) error {
	return nil
}

func validateColor(color []string) error {

	return nil
}

func (u *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Product{}, err
	}
	return u, nil
}

func (u *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (u *Product) FindProductByID(db *gorm.DB, uid uint32) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}
	return u, nil
}

func (u *Product) UpdateAProduct(db *gorm.DB, uid uint32) (*Product, error) {
	db = db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"title":      u.Title,
			"desc":       u.Desc,
			"img":        u.Img,
			"price":      u.Price,
			"in_stock":   u.InStock,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Product{}, db.Error
	}
	err := db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Product{}, err
	}
	return u, nil
}

func (u *Product) DeleteAProduct(db *gorm.DB, uid uint32) (string, error) {
	db = db.Debug().Model(&Product{}).Where("id = ?", uid).Take(&Product{}).Delete(&Product{})
	if db.Error != nil {
		return "", db.Error
	}
	if db.RowsAffected > 0 {
		return "Product deleted", nil
	}
	return "Product not found", nil
}

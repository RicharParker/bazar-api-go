package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rchavez-dev/fullstack/api/models"
)

var users = []models.User{
	models.User{
		Username: "Harim Castellanos",
		Email:    "hca@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "Mary Martinez",
		Email:    "mary@gmail.com",
		Password: "password",
	},
}

var products = []models.Product{
	models.Product{
		Title: "Product 1",
		Price: 9.99,
	},
	models.Product{
		Title: "Product 2",
		Price: 19.99,
	},
}

var orders = []models.Order{
	models.Order{
		UserID:   "1",
		Products: []models.Product{products[0]},
		Amount:   9.99,
		Status:   true,
	},
	models.Order{
		UserID:   "2",
		Products: []models.Product{products[1]},
		Amount:   19.99,
		Status:   false,
	},
}

var carts = []models.Cart{
	models.Cart{
		UserID:   "1",
		Products: []models.Product{products[0]},
	},
	models.Cart{
		UserID:   "2",
		Products: []models.Product{products[1]},
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Product{}, &models.Order{}, &models.Cart{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Cart{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range products {
		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
	}

	for i := range orders {
		err = db.Debug().Model(&models.Order{}).Create(&orders[i]).Error
		if err != nil {
			log.Fatalf("cannot seed orders table: %v", err)
		}
	}

	for i := range carts {
		err = db.Debug().Model(&models.Cart{}).Create(&carts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed orders table: %v", err)
		}
	}
}

package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/luthfi/pos/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Menunggu Database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("GAgal konek ke database: ", err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Member{},
	)

	DB = db
	fmt.Println("Database berhasil terhubunng dan migrasi selesai!")

}

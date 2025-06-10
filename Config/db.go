package Config

import (
	"GRPC/model"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initconfig() {
	// Read environment variable to decide which config file to load
	env := os.Getenv("APP_ENV")
	fmt.Println("Env is ", env)
	if env == "docker" {
		viper.SetConfigFile("config.docker.yaml")
	} else {
		viper.SetConfigFile("config.yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func DbConnect() {
	initconfig()

	dsn := viper.GetString("dsn")
	fmt.Println("Dsn is ", dsn)
	if dsn == "" {
		log.Println("Config file doesn't have dsn...")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with database...%v", err)
	}
	fmt.Println("DB connection is successful...")

	// Auto-migrate your models
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Migration failed..%v", err)
	}
	err = db.AutoMigrate(&model.Farmer{})
	if err != nil {
		log.Fatalf("Migration failed..%v", err)
	}
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		log.Fatalf("Migration failed..%v", err)
	}
	err = db.AutoMigrate(&model.Commodity{})
	if err != nil {
		log.Fatalf("Migration failed..%v", err)
	}
	err = db.AutoMigrate(&model.Bidding{})
	if err != nil {
		log.Fatalf("Migration failed..%v", err)
	}

	DB = db
}

package Config

import (
	"GRPC/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initconfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error to read from configuration file...%s", err)
	}
}

func DbConnect() {
	initconfig()
	dsn := viper.GetString("dsn")
	if dsn == "" {
		log.Println("Config file doesn't have dsn...")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with database...%v", err)
	}
	fmt.Println("DB connection is succesfull...")

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

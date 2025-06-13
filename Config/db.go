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

func DbConnect() {
	env := os.Getenv("APP_ENV")
	fmt.Println("Env is ", env)

	var dsn string

	if env == "docker" {

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, name)
	} else {
		viper.SetConfigFile("Config/config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		dsn = viper.GetString("dsn")
	}

	fmt.Println("Dsn is ", dsn)
	if dsn == "" {
		log.Fatal("DSN is empty! Check your config or environment variables.")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with database...%v", err)
	}
	fmt.Println("DB connection is successful...")

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

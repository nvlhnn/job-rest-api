package config

import (
	"dansmultipro/recruitment/model"
	"errors"
	"fmt"
	"log"
	"os"

	// "github.com/ydhnwb/golang_api/entity"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func OpenConnection() *gorm.DB{


	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}		
	

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	fmt.Println("migrating...")
	db.AutoMigrate(&model.User{})

	if db.Migrator().HasTable(&model.User{}){
		if err := db.First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			user := model.User{
				Username: "nvlhnn",
				Password: hashAndSalt([]byte("11111111")),				
			}
		
			err := db.Create(&user).Error
			if err != nil {
				log.Println("[seeding error] : ",err )
			}
		}
	}


	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
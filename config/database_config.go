package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloGinAndGorm/entity"
	//"os"
)

// SetupDatabaseConnection method creates a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	//errEnv := godotenv.Load()
	//if errEnv != nil {
	//	panic("Failed to load env file")
	//}
	//
	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbName := os.Getenv("DB_NAME")
	dbUser := "testuser"
	dbPass := "testpassword"
	dbHost := "helloginandgorm_mysql_1"
	dbName := "hello_gin_and_gorm"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection to database")
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}
// CloseDatabaseConnection method closes a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}


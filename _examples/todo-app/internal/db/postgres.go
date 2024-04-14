package db

import (
	//"log"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

// DbConnection create database connection
func InitPostgres(masterDSN string) error {
	var db = DB
	logMode := viper.GetBool("DB_LOG_MODE")
	// debug := viper.GetBool("DEBUG")

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	fmt.Println("Trying to connect to database...")
	db, _ = gorm.Open(postgres.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})

	fmt.Println("Database connection successful.")

	DB = db
	return nil
}

/* GetPostgresDB get postgres database instance.
* @return gorm.DB instance.
 */
func GetPostgresDB() *gorm.DB {
	return DB
}

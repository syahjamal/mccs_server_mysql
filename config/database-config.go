package config

import (
	"github.com/cengsin/oracle"
	_ "github.com/godror/godror"
	"github.com/syahjamal/mccs_server_mysql/model"
	"gorm.io/gorm"
)

func SetupDatabaseConn() *gorm.DB {

	db, err := gorm.Open(oracle.Open(`user="MFG" password="mfg_lge" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=150.150.242.55)(PORT=1527))(CONNECT_DATA=(SERVICE_NAME=LSY)))"`), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection to database")
	} else {
		println("Connected to database")
	}

	db.AutoMigrate(&model.MccsUser{})
	return db
}

/**
//SetupDatabaseConn is creating new connection to our database mysql
func SetupDatabaseConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to create a connection to database")
	}

	db.AutoMigrate(&model.MccsUser{})
	return db
}**/

//CloseDatabaseConn method is closing
func CloseDatabaseConn(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSql.Close()
}

package db

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
	// Blank import of postgres dialect is required according to Gorm docs
)

var once sync.Once
var db *gorm.DB

var replica *gorm.DB

func dbURL() string {
	var sslmode string
	sslmode = "sslmode=disable"
	dbPort, _ := strconv.Atoi(os.Getenv("DBPORT"))
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s %s",
		os.Getenv("DBHOST"),
		dbPort,
		os.Getenv("DBUSER"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPASSWORD"),
		sslmode,
	)
}

//Get awlays returns the primary database
func Get() *gorm.DB {
	once.Do(func() {
		db1, err := gorm.Open("postgres", dbURL())
		if err != nil {
			panic(err.Error())
		}
		dbase := db1.DB()
		err = dbase.Ping()
		if err != nil {
			panic(err.Error())
		}
		db = db1
	})
	return db
}

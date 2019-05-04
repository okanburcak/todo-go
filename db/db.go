package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/naughtydevelopment/todo-go/config"
	"log"
)

var db *gorm.DB
var err error

func Init() {
	c := config.GetConfig()

	user := c.GetString("db.user")
	pass := c.GetString("db.pass")
	host := c.GetString("db.host")
	port := c.GetString("db.port")
	name := c.GetString("db.name")

	dbi := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, pass, host, port, name)

	db, err = gorm.Open("postgres", dbi)

	if err != nil {
		log.Println("DB connection failed.", err.Error())
	}

	log.Println("DB connection success.")

}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

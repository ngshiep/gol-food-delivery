package main

import (
	"awesomeProject/module/restaurant/transport/ginrestuarant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)

	r := gin.Default()

	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")
	restaurants.POST("", ginrestuarant.CreateRestaurant(db))

	r.Run()
}

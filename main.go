package main

import (
	"awesomeProject/component/appctx"
	"awesomeProject/middleware"
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
	db = db.Debug()
	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	restaurants.GET("", ginrestuarant.ListRestaurant(appContext))
	restaurants.POST("", ginrestuarant.CreateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestuarant.DeleteRestaurant(appContext))

	r.Run()
}

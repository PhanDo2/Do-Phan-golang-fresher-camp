package main

import (
	"FoodDelivery/component"
	"FoodDelivery/middleware"
	"FoodDelivery/modules/restaurant/restauranttransport/ginrestaurant"
	"FoodDelivery/modules/user/usertransport/ginuser"
	"FoodDelivery/pubsub"
	"FoodDelivery/pubsub/pblocal"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionString")

	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db, secretKey, pblocal.NewPubSub()); err != nil {
		log.Fatalln(err)
	}

}
func runService(db *gorm.DB, secretKey string, pb pubsub.Pubsub) error {
	appCtx := component.NewAppContext(db, secretKey, pb)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", ginuser.Register(appCtx))
	r.POST("/login", ginuser.Login(appCtx))
	r.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := r.Group("/restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
	}

	return r.Run()
}

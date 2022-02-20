package main

import (
	"Fresher_go/component"
	"Fresher_go/component/uploadprovider"
	"Fresher_go/middleware"
	"Fresher_go/modules/notes/notetransport/ginnote"
	"Fresher_go/modules/restaurents/restauranttransport/ginrestaurants"
	"Fresher_go/modules/upload/uploadtransport/ginupload"
	"Fresher_go/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

//CREATE TABLE `note` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`name` varchar(50) NOT NULL,
//`addr` varchar(255) NOT NULL,
//`status` int(11) NOT NULL DEFAULT '1',
//PRIMARY KEY (`id`),
//KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8
type Note struct {
	Id     int    `json:"id,omitempty" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"addr" gorm:"column:addr"`
	Status int    `json:"status" gorm:"column:status"`
}

func (Note) TableName() string {
	return "note"
}

func main() {
	dsn := os.Getenv("DBConnection")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	//dsn := "root:Do@19012000@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	//db.AutoMigrate()
	if err != runService(db, s3Provider) {
		log.Fatalln(err)
	}
}

//CRUD

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	n := gin.Default()
	n.Use(middleware.Recover(appCtx))
	n.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	n.POST("/upload", ginupload.Upload(appCtx))
	n.POST("/register", ginuser.Register(appCtx))
	notes := n.Group("/notes")
	{
		notes.POST("", ginnote.CreateNote(appCtx))
		notes.GET("", ginnote.ListNote(appCtx))
		notes.GET("/:id", ginnote.GetNote(appCtx))
		notes.DELETE("/:id", ginnote.DeleteNote(appCtx))
		notes.POST("/:id", ginnote.UpdateNote(appCtx))

	}
	restaurants := n.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurants.CreateRetaurents(appCtx))
		restaurants.POST("/:id", ginrestaurants.UpdateRetaurents(appCtx))
		restaurants.GET("", ginrestaurants.ListRetaurents(appCtx))
		restaurants.DELETE("/:id", ginrestaurants.DeleteRetaurents(appCtx))
		restaurants.GET("/:id", ginrestaurants.GetRetaurents(appCtx))
	}

	return n.Run()
}

package main

import (
	"Fresher_go/component"
	"Fresher_go/modules/notes/notetransport/ginnote"
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

	//dsn := "root:Do@19012000@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DBConnection")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err != runService(db) {
		log.Fatalln(err)
	}
}

//CRUD

func runService(db *gorm.DB) error {
	n := gin.Default()
	n.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	appCtx := component.NewAppConntect(db)
	notes := n.Group("/notes")
	{
		notes.POST("", ginnote.CreateNote(appCtx))
		notes.GET("", ginnote.ListNote(appCtx))
		notes.GET("/:id", ginnote.GetNote(appCtx))
		notes.DELETE("/:id", ginnote.DeleteNote(appCtx))
		notes.POST("/:id", ginnote.UpdateNote(appCtx))
	}

	return n.Run()
}

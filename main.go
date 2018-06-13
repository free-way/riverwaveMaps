package main

import(
	"gopkg.in/ini.v1"
	"flag"
	"fmt"
	"github.com/free-way/riverwaveMaps/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/free-way/riverwaveMaps/models"
	"github.com/gin-gonic/gin"
	"github.com/free-way/riverwaveMaps/handlers"
)
var(
	err error
	cfg *ini.File
)

func init(){
	cfgFlag := flag.String("config", "", "Env file")
	flag.Parse()
	cfg,err = ini.Load(*cfgFlag)
	if err != nil {
		fmt.Println("could not load configuration file due to: ", err.Error())
		panic(err)
	}

	helpers.DB,err = gorm.Open("mysql",cfg.Section("Database").Key("ConnectionString").String())
	if err != nil{
		fmt.Println("Could not connect to the database due to", err.Error())
	}

	models.RunMigration()

}
/*
	TODO: Add Middleware to validate user Access
	TODO: Add Middleware to call events microservice and validate if the event id passed exists
 */

func main(){
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200,"Welcome!")
	})
	r.POST("/events/:event/maps",handlers.AddMap)
	r.GET("/events/:event/maps",handlers.GetMaps)
	r.DELETE("/events/:event/maps/:map",handlers.DeleteMap)


	r.Run(cfg.Section("Server").Key("Port").String())
}
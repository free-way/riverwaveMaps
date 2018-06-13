package main

import(
	"gopkg.in/ini.v1"
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
	"github.com/free-way/riverwaveMaps/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/free-way/riverwaveMaps/models"
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

func main(){
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200,"Welcome!")
	})

	r.Run(cfg.Section("Server").Key("Port").String())
}
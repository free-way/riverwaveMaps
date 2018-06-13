package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/free-way/riverwaveMaps/models"
	"strconv"
	"github.com/free-way/riverwaveMaps/helpers"
	"net/http"
)

var(
	err error
)

func AddMap(ctx *gin.Context)  {
	var payload models.Map
	eventId,_ := strconv.Atoi(ctx.Param("event"))
	ctx.BindJSON(&payload)
	payload.EventId = eventId
	if err = payload.Validate(); err != nil{
		//return 400 validation error
		ctx.JSON(400,err.Error())
		ctx.Abort()
		return
	}

	helpers.DB.Save(&payload)
	ctx.JSON(200,payload)
}

func GetMaps(ctx *gin.Context){
	var res []models.Map
	eventId,_ := strconv.Atoi(ctx.Param("event"))
	helpers.DB.Where("event_id = ?",eventId).Find(&res)
	ctx.JSON(200,res)
}

func DeleteMap(ctx *gin.Context){
	var mapInstance models.Map
	eventId,_ := strconv.Atoi(ctx.Param("event"))
	mapId,_ := strconv.Atoi(ctx.Param("map"))
	if helpers.DB.Where("id = ? && event_id = ?",mapId,eventId).Find(&mapInstance).RecordNotFound(){
		ctx.JSON(http.StatusNotFound,map[string]string{
			"Message":"Map Not Found",
		})
		ctx.Abort()
		return
	}
	helpers.DB.Delete(&mapInstance)
	ctx.JSON(200,map[string]string{
		"Message":"Map Successfully Deleted!",
	})
}

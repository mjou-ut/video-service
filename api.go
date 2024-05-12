package video_service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/mjou-ut/video-service/db"
	video "github.com/mjou-ut/video-service/models"
)

func getVideo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	v, err := db.GetVideo(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, v)
}

func getVideos(ctx *gin.Context) {
	vs := db.GetVideos()
	ctx.JSON(http.StatusOK, vs)
}

func addVideo(ctx *gin.Context) {
	v := video.NewVideo()
	db.SaveVideo(v)
	ctx.JSON(http.StatusOK, v)
}

func addVideos(ctx *gin.Context) {
	paramNum := ctx.Params.ByName("n")
	num, err := strconv.Atoi(paramNum)
	if err != nil || num < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid number",
		})
		return
	}

	for i := 0; i < num; i++ {
		v := video.NewVideo()
		db.SaveVideo(v)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d videos added", num),
	})
}

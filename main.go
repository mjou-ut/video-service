package video_service

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/videos", getVideos)
	r.GET("/videos/:id", getVideo)
	r.POST("/videos/:n", addVideos)
	r.POST("/videos", addVideo)
	r.Run("0.0.0.0:3000")
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

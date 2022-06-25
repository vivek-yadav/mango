package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/mango/config"
)

func StartGin(conf config.Config) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprint(":", conf.Serve.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

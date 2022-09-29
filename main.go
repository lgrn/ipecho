package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// do not trust proxies to provide "real ip" via header
	r.SetTrustedProxies(nil)

	r.GET("/json", func(c *gin.Context) {
		// c.Request.Header returns []string, use 'strings' to convert to string
		user_agent := strings.Join(c.Request.Header["User-Agent"], "")
		c.JSON(http.StatusOK, gin.H{
			"IP":        c.ClientIP(),
			"USERAGENT": user_agent,
		})
	})

	r.GET("/text", func(c *gin.Context) {
		ip := c.ClientIP() + "\n"
		c.String(http.StatusOK, ip)
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(400, "HTTP Error 400: Bad Request.\nSupported endpoints are /json and /text\n")
	})

	// port is always 80, binary must run as root
	err := r.Run("0.0.0.0:80")
	if err != nil {
		fmt.Println(err)
	}
}

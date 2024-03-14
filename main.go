package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type stringSlice []string

func (i *stringSlice) String() string {
	return strings.Join(*i, ",")
}

func (i *stringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var trustedProxies stringSlice
	flag.Var(&trustedProxies, "proxy", "IP address of reverse proxy that will contact us, and we trust to provide correct headers (this flag can be repeated)")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if len(trustedProxies) > 0 {
		if err := r.SetTrustedProxies(trustedProxies); err != nil {
			log.Fatalf("Tried and failed to set trusted proxies: %v", err)
		}
	} else {
		r.SetTrustedProxies(nil)
	}

	r.GET("/json", func(c *gin.Context) {
		// c.Request.Header returns []string, use 'strings' to convert to string
		user_agent := strings.Join(c.Request.Header["User-Agent"], "")
		c.JSON(http.StatusOK, gin.H{
			"IP":        c.ClientIP(),
			"USERAGENT": user_agent,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, c.ClientIP()+"\n")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(400, "HTTP Error 400: Bad Request.\nSupported endpoints are / and /json\n")
	})

	// port is always 80, binary must run as root
	err := r.Run("0.0.0.0:80")
	if err != nil {
		fmt.Println(err)
	}
}

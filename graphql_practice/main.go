package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	host := "localhost"
	port := "7777"

	pathQL := "/graphql"
	r := gin.Default()

	//route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	log.Println(host + ":" + port + pathQL + "에서 서버를 실행중입니다.")
	log.Fatalln(r.Run(host + ":" + port))
}

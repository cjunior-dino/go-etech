package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func fatorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/fatorial", func(ctx *gin.Context) {
		num, _ := strconv.Atoi(ctx.Query("num"))
		result := fatorial(num)
		ctx.JSON(http.StatusOK, gin.H{"result": result})
	})

	router.Run(":8080")
}

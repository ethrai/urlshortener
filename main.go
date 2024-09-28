package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

const ctxStoreKey = "dbConn"

func main() {
	db := NewStore(os.Getenv("DB_URL"))

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.Use(func(ctx *gin.Context) {
		ctx.Set(ctxStoreKey, db)
		ctx.Next()
	})

	e.GET("/", homeHandler)
	e.POST("/", createURLHandler)
	e.GET("/:alias", redirectHandler)

	e.Run("127.0.0.1:8080")
}

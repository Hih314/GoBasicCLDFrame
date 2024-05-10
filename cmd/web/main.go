package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个中间件函数
func MidTest(ctx *gin.Context) {
	fmt.Println("全局中间件 => ", ctx.FullPath())
}

func main() {
	app := gin.Default()

	app.Use(MidTest)
	app.GET("/", func(ctx *gin.Context) {
		fmt.Println("中间件执行1")
	}, func(ctx *gin.Context) {
		fmt.Println("中间件执行2")
	}, func(r *gin.Context) {
		r.JSON(200, gin.H{
			"msg": "success",
		})
	})

	v1 := app.Group("/api")
	v1.Use(MidTest)
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "api ok",
		})
	})

	v2 := v1.Group("/api2")
	v2.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	app.Run(":3000")
}

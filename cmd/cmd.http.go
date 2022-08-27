package cmd

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go-template/app"
	"go-template/configs"
	"go-template/services"
	"go.elastic.co/apm/module/apmgin"
	"net/http"
	"time"
)

func HTTPConfigs() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "GET", "POST", "DELETE", "OPTION"},
		AllowHeaders: []string{"Origin",
			"Access-Control-Allow-Origin",
			"Content-Type", "Content-Length", "Access-Control-Allow-Methods", "x-access-token", "Accept-Encoding", "X-CSRF-Token", "Authorization", "x-captcha-token", "x-api-duration", "token"},
		ExposeHeaders:    []string{"x-access-token", "x-api-duration", "x-captcha-token"},
		AllowWildcard:    true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func HTTPServe(ctx *cli.Context) error {
	serve := gin.New()
	serve.RedirectTrailingSlash = false
	serve.Use(gin.Recovery(), HTTPConfigs(), gin.Logger(), func(c *gin.Context) {
		//for k, v := range utils.HeaderDefault {
		//	c.Writer.Header().Set(k, v)
		//}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}, apmgin.Middleware(serve))

	SetupRouting(serve)

	return serve.Run(":" + configs.GlobalConf.HTTPPort)
}

func SetupRouting(serve *gin.Engine) {
	// Ping
	serve.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	service := services.Services{App: app.App}
	groupApi := serve.Group("/viettel")
	groupApi.POST("/merchantGetCodeV2", service.GetCodeController)
	groupApi.POST("/holdCodeWithTransaction", service.HoldCodeController)
}

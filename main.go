package main

import (
	"github.com/kainnsoft/tmanager/internal/routes"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")

	routes.InitializeRoutes(router)

	// router.GET("/", func(ctx *gin.Context) {
	// 	// Call the HTML method of the Context to render a template
	// 	ctx.HTML(
	// 		// Set the HTTP status to 200 (OK)
	// 		http.StatusOK,
	// 		// Use the index.html template
	// 		"index.html",
	// 		// Pass the data that the page uses
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	// Start serving the application
	router.Run(":9999")
}

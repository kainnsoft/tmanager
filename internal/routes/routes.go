package routes

import (
	"github.com/kainnsoft/tmanager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	// определение роута главной страницы
	router.GET("/", handlers.ShowIndexPage)

	// Обработчик GET-запросов на /article/view/некоторый_article_id
	router.GET("/article/view/:article_id", handlers.GetArticle)
}

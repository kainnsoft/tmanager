package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kainnsoft/tmanager/internal/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(ctx *gin.Context) {
	articles := models.GetAllArticles()

	srt := ctx.Request.Header.Get("Accept")
	fmt.Println("Heder = ", srt)

	// Call the HTML method of the Context to render a template
	ctx.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}

func GetArticle(ctx *gin.Context) {
	// Проверим валидность ID
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		// Проверим существование топика
		if article, err := models.GetArticleByID(articleID); err == nil {
			// Вызовем метод HTML из Контекста Gin для обработки шаблона
			ctx.HTML(
				// Зададим HTTP статус 200 (OK)
				http.StatusOK,
				// Используем шаблон article.html
				"article.html",
				// Передадим данные в шаблон
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// Если топика нет, прервём с ошибкой
			ctx.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// При некорректном ID в URL, прервём с ошибкой
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

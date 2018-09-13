package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/ryomak/go-vue-example/middleware"
)

func AssignAPI(g *echo.Group, db *gorm.DB) {
	g.Use(middleware.TransactionMiddleware(db))
	g.GET("/article", getAllArticleHandler)
	g.GET("/article/:id", getArticleHandler)

}

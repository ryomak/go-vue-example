package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func getArticleHandler(c echo.Context) error {
	c.String(http.StatusOK, "fdfd")
	return nil
}

func getAllArticleHandler(c echo.Context) error {
	return nil
}

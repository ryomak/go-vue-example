package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetDB(c echo.Context) *gorm.DB {
	return c.Get("DB").(*gorm.DB)
}

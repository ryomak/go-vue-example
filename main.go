package main

import (
	"github.com/labstack/echo"
	"github.com/ryomak/go-vue-example/config"
	"github.com/ryomak/go-vue-example/controllers"
	"github.com/labstack/echo/middleware"
)

func init() {
	if err := config.LoadConfig("config.toml"); err != nil {
		panic(err)
	}
}

func main() {
	conf := config.GetConfig()
	db := config.InitDB(conf.DBConfig)

	e := echo.New()
	e.Use(middleware.Logger())
	e.File("/","public/index.html")
	e.Static("/","public")

	controllers.AssignAPI(e.Group("/api/v1"), db)

	e.Start(":" + conf.Port)
}

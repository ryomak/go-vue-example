package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ryomak/go-vue-example/config"
	"github.com/ryomak/go-vue-example/controllers"
)

func init() {
	filenPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	if err := config.LoadConfig(*filenPath); err != nil {
		panic(err)
	}
}

func main() {
	conf := config.GetConfig()
	db := config.InitDB(conf.DBConfig)

	e := echo.New()
	e.Use(middleware.Logger())
	e.File("/", "public/index.html")
	e.Static("/", "public")

	controllers.AssignAPI(e.Group("/api/v1"), db)

	e.Start(":" + conf.Port)
}

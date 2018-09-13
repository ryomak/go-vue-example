package main

import (
	"github.com/labstack/echo"
	"github.com/ryomak/go-vue-example/config"
	"github.com/ryomak/go-vue-example/controllers"
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
	e.File("/","public/index.html")
	controllers.AssignAPI(e.Group("/api/v1"), db)
	e.Start(":" + conf.Port)
}

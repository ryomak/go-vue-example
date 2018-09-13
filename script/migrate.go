package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/ryomak/go-vue-example/config"
	"github.com/ryomak/go-vue-example/models"
)

func main() {
	filenPath := flag.String("config", "./config.toml", "path to config file")
	flag.Parse()
	var conf struct {
		DBConfig config.DBConfig `toml:"database"`
	}
	if _, err := toml.DecodeFile(*filenPath, &conf); err != nil {
		panic(err)
	}
	db := config.InitDB(conf.DBConfig)
	if err := models.Migrate(db); err != nil {
		panic(err)
	}
}

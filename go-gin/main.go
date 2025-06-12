package main

import (
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func main() {
	cfg := loadEnv()

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
				"year":  time.Now().Year(),
			},
		)
	})

	router.Run()
}
type config struct {
	DATABASE_USER     string `env:"DATABASE_USER" envDefault:"postgres"`
	DATABASE_PASSWORD string `env:"DATABASE_PASSWORD" envDefault:"password"`
	DATABASE_HOST     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DATABASE_PORT     string `env:"DATABASE_PORT" envDefault:"5432"`
	DATABASE_NAME     string `env:"DATABASE_NAME" envDefault:"database"`
}

func loadEnv() config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	var cfg config
	cfg, err = env.ParseAs[config]()
	if err != nil {
		panic("failed to get environment variables")
	}

	return cfg
}

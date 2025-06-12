package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lwandokasuba/project-management/ent"
	_ "github.com/lib/pq"
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

func connectDB(cfg config) *ent.Client {
	client, err := ent.Open("postgres", "host="+cfg.DATABASE_HOST+" user="+cfg.DATABASE_USER+
		" password="+cfg.DATABASE_PASSWORD+" dbname="+cfg.DATABASE_NAME+
		" port="+cfg.DATABASE_PORT+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
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

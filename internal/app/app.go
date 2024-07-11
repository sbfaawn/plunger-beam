package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plunger-beam/internal/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func Run() {
	server := gin.Default()

	cfg := readConfig()
	fmt.Println(cfg)

	group := server.Group("/api/v1")
	group.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"status":  "Success",
			"data":    "",
			"message": "",
		})
	})

	server.Run(":" + cfg.Server.Port)
}

func readConfig() config.Config {
	env := flag.String("env", "development", "Environment (development|production)")
	flag.Parse()

	configPath := filepath.Join("configs", "config."+*env+".yaml")
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	return cfg
}

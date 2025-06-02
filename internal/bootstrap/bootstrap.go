package bootstrap

import (
	"fmt"
	"log"
	"os"

	"github.com/adityarifqyfauzan/go-boilerplate/cmd"
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/container"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

func Init() {
	if err := cmd.RefreshRepositoryContainer(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.RefreshModuleContainer(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load env file: %v", err)
	}

	dig := dig.New()
	container.BuildContainer(dig)
	defer config.CloseDB()

	app := &cli.App{
		Name: "Go Boilerplate Starter Kit",
		Commands: []*cli.Command{
			cmd.ModelCommand,
			cmd.MigrationCommand,
			cmd.MakeMigrationCommand,
			cmd.MigrateAllCommand,
			cmd.MigrateFreshCommand,
			cmd.RollbackAllCommand,
			cmd.RollbackBatchCommand,
			cmd.RollbackCommand,
		},
	}

	if len(os.Args) > 1 {
		err := app.Run(os.Args)
		if err != nil {
			log.Default().Fatalf("failed to run app: %v", err)
		}
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

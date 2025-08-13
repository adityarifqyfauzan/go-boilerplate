package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/routes"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/middleware"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/opentelemetry"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func RestAPI(ctx context.Context, conf *config.Config) {
	if helper.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(middleware.I18nMiddleware())
	r.Use(otelgin.Middleware(opentelemetry.GetServiceName()))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		i18n := translator.NewTranslator(c.Value("localizer").(*i18n.Localizer))
		c.JSON(200, gin.H{
			"message": i18n.T("hello", map[string]any{
				"Name": "Aditya",
			}),
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	routes.Init(r, conf)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler:     r,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	// start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

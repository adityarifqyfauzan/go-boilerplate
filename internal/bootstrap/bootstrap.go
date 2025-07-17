package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/command"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/routes"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/middleware"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/opentelemetry"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Init() {
	if err := command.RefreshRepositoryContainer(); err != nil {
		log.Fatal(err)
	}

	if err := command.RefreshModuleContainer(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load env file: %v", err)
	}

	routerConfig := config.New()
	defer config.CloseDB()

	translator.Init("locales")

	cmd := command.App
	if len(os.Args) > 1 {
		err := cmd.Run(os.Args)
		if err != nil {
			log.Default().Fatalf("failed to run app: %v", err)
		}
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// setup OpenTelemetry
	otelShutdown, err := opentelemetry.SetupOTelSDK(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	r := gin.Default()
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

	routes.Init(r, routerConfig)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler:      r,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
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

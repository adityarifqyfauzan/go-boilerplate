package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

var CreateModuleCommand = &cli.Command{
	Name:  "make:module",
	Usage: "Create a new module",
	Action: func(c *cli.Context) error {
		name := c.Args().Get(0)
		fmt.Println("🚀 Generate module:", name)

		err := NewModuleGenerator(name).Generate()
		if err != nil {
			return fmt.Errorf("failed to generate module: %w", err)
		}

		return nil
	},
}

var CreateWorkerModuleCommand = &cli.Command{
	Name:  "make:worker",
	Usage: "Create a new worker module",
	Action: func(c *cli.Context) error {
		name := c.Args().Get(0)
		fmt.Println("🚀 Generate worker module:", name)

		err := NewModuleGenerator(name).GenerateWorker()
		if err != nil {
			return fmt.Errorf("failed to generate worker module: %w", err)
		}

		return nil
	},
}

type ModuleGenerator struct {
	ModuleName string
}

func NewModuleGenerator(moduleName string) *ModuleGenerator {
	return &ModuleGenerator{
		ModuleName: strings.ToLower(moduleName),
	}
}

func (g *ModuleGenerator) GenerateWorker() error {
	modulePath := filepath.Join("internal", "module", g.ModuleName)

	// check if module already exists
	if _, err := os.Stat(modulePath); !os.IsNotExist(err) {
		return fmt.Errorf("module already exists")
	}

	if err := os.MkdirAll(modulePath, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	files := map[string]string{
		"model.go":            modelTemplate,
		"dto.go":              workerDto,
		"service.go":          workerService,
		"worker.go":           worker,
		"local_repository.go": repositoryTemplate,
	}

	for filename, content := range files {
		filePath := filepath.Join(modulePath, filename)
		tmpl, err := template.New(filename).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template for %s: %w", filename, err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filename, err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, g); err != nil {
			return fmt.Errorf("failed to execute template for %s: %w", filename, err)
		}
	}

	fmt.Printf("Module '%s' generated successfully in %s\n", g.ModuleName, modulePath)
	return nil
}

func (g *ModuleGenerator) Generate() error {
	modulePath := filepath.Join("internal", "module", g.ModuleName)

	// check if module already exists
	if _, err := os.Stat(modulePath); !os.IsNotExist(err) {
		return fmt.Errorf("module already exists")
	}

	if err := os.MkdirAll(modulePath, 0755); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	files := map[string]string{
		"model.go":            modelTemplate,
		"dto.go":              dtoTemplate,
		"handler.go":          handlerTemplate,
		"service.go":          serviceTemplate,
		"route.go":            routeTemplate,
		"local_repository.go": repositoryTemplate,
	}

	for filename, content := range files {
		filePath := filepath.Join(modulePath, filename)
		tmpl, err := template.New(filename).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template for %s: %w", filename, err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filename, err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, g); err != nil {
			return fmt.Errorf("failed to execute template for %s: %w", filename, err)
		}
	}

	fmt.Printf("Module '%s' generated successfully in %s\n", g.ModuleName, modulePath)
	return nil
}

const modelTemplate = `package {{.ModuleName}}
`

const dtoTemplate = `package {{.ModuleName}}
`

const handlerTemplate = `package {{.ModuleName}}

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(
	service Service,
) handler {
	return handler{
		service: service,
	}
}

func (h *handler) Get(c *gin.Context) {
	// Implement get handler
}

`

const serviceTemplate = `package {{.ModuleName}}

import (
	"context"
	"net/http"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"gorm.io/gorm"
)

type Service interface {
	Get(ctx context.Context, id int) *helper.ApiResponse
}

type service struct {
	localRepository LocalRepository
	db *gorm.DB
}

func NewService(
	db *gorm.DB,
	localRepository LocalRepository,
) Service {
	return &service{
		db: db,
		localRepository: localRepository,
	}
}

func (s *service) Get(ctx context.Context, id int) *helper.ApiResponse {
	// Implement get service
	return helper.NewApiResponse(http.StatusOK, "success", nil)
}
`

const routeTemplate = `package {{.ModuleName}}

import (
	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.RouterGroup, config *config.Config) {

	service := NewService(
		config.DB,
		NewLocalRepository(config.DB),
	)

	handler := NewHandler(service)

	{{.ModuleName}}Route := route.Group("{{.ModuleName}}")
	{{.ModuleName}}Route.GET("/get", handler.Get)
}
`

const repositoryTemplate = `package {{.ModuleName}}

import (
	"gorm.io/gorm"
)

type LocalRepository interface {
}

type localRepository struct {
	db *gorm.DB
}

func NewLocalRepository(
	db *gorm.DB,
) LocalRepository {
	return &localRepository{
		db: db,
	}
}
`

const worker = `package {{.ModuleName}}

import (
	"context"
	"encoding/json"
	"log"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

func ExampleWorker() func(ctx context.Context, ch *amqp.Channel, conf *config.Config) {
	return func(ctx context.Context, ch *amqp.Channel, conf *config.Config) {
		service := NewService(
			config.DB,
			NewLocalRepository(config.DB),
		)

		message, err := rabbitmq.Consume(ch, &rabbitmq.PublishOption{
			Topic: "example",
		})
		if err != nil {
			panic(err)
		}

		log.Println("example worker started")

		for msg := range message {
			tr := otel.Tracer("example-worker")
			ctx, span := tr.Start(ctx, "ExampleWorker")

			log.Println("example worker received message")

			var request ExampleRequest
			if err := json.Unmarshal(msg.Body, &request); err != nil {
				panic(err)
			}

			if err := service.Example(ctx, request.Name); err != nil {
				log.Printf("failed to example: %v", err)
			}

			span.End()

			log.Println("example worker processed message")
		}
	}
}
`

const workerService = `package {{.ModuleName}}

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	Example(ctx context.Context, name string) error
}

type service struct {
	localRepository LocalRepository
	db              *gorm.DB
}

func NewService(
	db *gorm.DB,
	localRepository LocalRepository,
) Service {
	return &service{
		db:              db,
		localRepository: localRepository,
	}
}

func (s *service) Example(ctx context.Context, name string) error {
	log.Println("Hello, ", name)
	return nil
}

`

const workerDto = `package {{.ModuleName}}

type ExampleRequest struct {
	Name string ` + "`" + `json:"name"` + "`" + `
}`

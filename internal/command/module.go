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
		fmt.Println("🚀 Generate model:", name)

		err := NewModuleGenerator(name).Generate()
		if err != nil {
			return fmt.Errorf("failed to generate module: %w", err)
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
		"container.go":        containerTemplate,
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

const containerTemplate = `package {{.ModuleName}}

import "go.uber.org/dig"

func InitContainer(container *dig.Container) {
	if err := container.Provide(NewLocalRepository); err != nil {
		panic(err)
	}

	if err := container.Provide(NewService); err != nil {
		panic(err)
	}
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

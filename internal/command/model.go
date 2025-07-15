package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/urfave/cli/v2"
)

var ModelCommand = &cli.Command{
	Name:  "make:model",
	Usage: "Generate model",
	Action: func(c *cli.Context) error {

		name := c.Args().Get(0)
		fmt.Println("🚀 Generate model:", name)

		err := generateModel(name)
		if err != nil {
			return err
		}

		err = RefreshRepositoryContainer()
		if err != nil {
			return err
		}

		return nil
	},
}

func generateModel(name string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	if name == "" {
		return fmt.Errorf("name is required")
	}

	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ToLower(name)

	// check if file already exists
	if _, err := os.Stat(wd + "/internal/model/" + name + ".go"); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists")
	}

	pattern := `package model

type %s struct {
	BaseModel
}

func (%s) TableName() string {
	return "%s"
}
	`
	model := strings.Split(name, "_")
	for _, m := range model {
		m = strings.Title(m)
		pattern = strings.ReplaceAll(pattern, m, m)
	}

	modelName := helper.ToPascalCase(name)

	content := fmt.Sprintf(pattern, modelName, modelName, helper.ToPlural(name))
	err = os.WriteFile(wd+"/internal/model/"+name+".go", []byte(content), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	return nil
}

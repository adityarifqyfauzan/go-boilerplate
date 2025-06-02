package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

var ModulCommand = &cli.Command{
	Name:  "make:modul",
	Usage: "Generate modul",
	Action: func(c *cli.Context) error {
		name := c.Args().Get(0)
		fmt.Println("🚀 Generate modul:", name)
		return nil
	},
}

func GenerateModul(name string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	moduleDir := fmt.Sprintf("%s/internal/module/", wd)

	if name == "" {
		return fmt.Errorf("name is required")
	}

	moduleName := strings.ToLower(removeSpecialCharsAndSpaces(name))

	// check if dir already exists
	if _, err := os.Stat(moduleDir + moduleName); !os.IsNotExist(err) {
		return fmt.Errorf("module already exists")
	}

	// create new dir
	if err := os.MkdirAll(moduleDir+moduleName, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create module directory: %w", err)
	}

	return nil
}

func removeSpecialCharsAndSpaces(input string) string {
	// Regex buat hapus semua karakter non-alfanumerik (A-Z, a-z, 0-9)
	// \W artinya non-word character (sama kayak [^a-zA-Z0-9_])
	// Kalau gak mau underscore (_) ikut, tinggal pakai [^a-zA-Z0-9]
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// Replace semua yang cocok dengan regex jadi string kosong
	cleaned := re.ReplaceAllString(input, "")

	return cleaned
}

package command

import (
	"log"
	"reflect"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/database/seeders"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var SeederCommand = &cli.Command{
	Name:  "seeder",
	Usage: "Run seeders",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "only",
			Aliases: []string{"o"},
			Usage:   "Run only one seeder",
		},
	},
	Action: func(c *cli.Context) error {
		if !helper.ConfirmProductionAction() {
			return nil
		}

		seeders.RegisterSeeders()
		RunAll(config.DB, c.String("only"))
		return nil
	},
}

func RunAll(db *gorm.DB, only string) {
	selectedSeeders := seeders.GetSeeders()

	if only != "" {
		selectedSeeders = []seeders.Seeder{}
		for _, s := range seeders.GetSeeders() {
			if reflect.TypeOf(s).Name() == only {
				selectedSeeders = append(selectedSeeders, s)
			}
		}
		if len(selectedSeeders) == 0 {
			log.Printf("Seeder %s not found.\n", only)
			return
		}
	}

	bar := progressbar.Default(int64(len(selectedSeeders)))
	successSeeder := []string{}
	for _, s := range selectedSeeders {
		err := db.Transaction(func(tx *gorm.DB) error {
			return s.Run(tx)
		})
		if err != nil {
			log.Printf("Seeder failed: %v\n", err)
			continue
		}

		successSeeder = append(successSeeder, reflect.TypeOf(s).Name())
		bar.Add(1)
	}

	if only != "" {
		log.Printf("✅ Seeder %s executed successfully\n", only)
		return
	}

	for _, v := range successSeeder {
		log.Printf("✅ Seeder %s executed successfully\n", v)
	}

	log.Println("✅ All seeders executed successfully.")
}

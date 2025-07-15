package command

import "github.com/urfave/cli/v2"

// register all command here 👇
var (
	App = &cli.App{
		Name: "Go Boilerplate Starter Kit",
		Commands: []*cli.Command{
			ModelCommand,
			CreateModuleCommand,
			MigrateCreateCommand,
			MigrateUpCommand,
			MigrateStatusCommand,
			MigrateRefreshCommand,
			MigrateDownToCommand,
			MigrateUpToCommand,
			MigrateDownCommand,
			SeederCommand,
		},
	}
)

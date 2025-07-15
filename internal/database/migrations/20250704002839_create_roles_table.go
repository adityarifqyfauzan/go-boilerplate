package migrations

import (
	"context"
	"database/sql"

	"github.com/ahmadfaizk/schema"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateRolesTable, downCreateRolesTable)
}

func upCreateRolesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return schema.Create(ctx, tx, "roles", func(table *schema.Blueprint) {
		table.UnsignedBigInteger("id", true).Primary()
		table.String("name", 50).Unique()
		table.String("slug", 50).Unique()
		table.Boolean("is_active").Default("TRUE")
		table.Timestamp("created_at").Default("CURRENT_TIMESTAMP")
		table.Timestamp("updated_at").Default("CURRENT_TIMESTAMP").UseCurrentOnUpdate()
	})
}

func downCreateRolesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return schema.Drop(ctx, tx, "roles")
}

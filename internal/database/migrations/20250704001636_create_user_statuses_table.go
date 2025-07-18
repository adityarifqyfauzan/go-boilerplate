package migrations

import (
	"context"
	"database/sql"

	"github.com/ahmadfaizk/schema"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserStatusesTable, downUserStatusesTable)
}

func upUserStatusesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return schema.Create(ctx, tx, "user_statuses", func(table *schema.Blueprint) {
		table.ID()
		table.String("name", 50).Unique()
		table.String("slug", 50).Unique()
		table.Timestamp("created_at").Default("CURRENT_TIMESTAMP")
		table.Timestamp("updated_at").Default("CURRENT_TIMESTAMP").UseCurrentOnUpdate()
	})
}

func downUserStatusesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return schema.Drop(ctx, tx, "user_statuses")
}

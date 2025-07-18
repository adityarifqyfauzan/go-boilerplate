package migrations

import (
	"context"
	"database/sql"

	"github.com/ahmadfaizk/schema"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserStatusHistoriesTable, downUserStatusHistoriesTable)
}

func upUserStatusHistoriesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return schema.Create(ctx, tx, "user_status_histories", func(table *schema.Blueprint) {
		table.ID()
		table.UnsignedBigInteger("user_id")
		table.UnsignedBigInteger("user_status_id")
		table.Integer("created_by").Default(0)
		table.Timestamp("created_at").Default("CURRENT_TIMESTAMP")
		table.Timestamp("updated_at").Default("CURRENT_TIMESTAMP").UseCurrentOnUpdate()
		table.Foreign("user_id").References("id").On("users")
		table.Foreign("user_status_id").References("id").On("user_statuses")
	})
}

func downUserStatusHistoriesTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return schema.Drop(ctx, tx, "user_status_histories")
}

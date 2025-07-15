package migrations

import (
	"context"
	"database/sql"

	"github.com/ahmadfaizk/schema"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return schema.Create(ctx, tx, "users", func(table *schema.Blueprint) {
		table.UnsignedBigInteger("id", true).Primary()
		table.String("name", 100)
		table.String("email", 255).Unique()
		table.String("password", 255)
		table.String("remember_token", 100).Nullable().Default("NULL")
		table.UnsignedBigInteger("user_status_id").Default(1)
		table.Timestamp("created_at").Default("CURRENT_TIMESTAMP")
		table.Timestamp("updated_at").Default("CURRENT_TIMESTAMP").UseCurrentOnUpdate()
		table.Timestamp("deleted_at").Nullable().Default("NULL").Index()

		table.Foreign("user_status_id").References("id").On("user_statuses")
	})
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return schema.Drop(ctx, tx, "users")
}

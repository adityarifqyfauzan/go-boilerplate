package migrations

import (
	"context"
	"database/sql"

	"github.com/ahmadfaizk/schema"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserDetailsTable, downUserDetailsTable)
}

func upUserDetailsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return schema.Create(ctx, tx, "user_details", func(table *schema.Blueprint) {
		table.UnsignedBigInteger("id", true).Primary()
		table.UnsignedBigInteger("user_id")
		table.Text("address").Nullable()
		table.String("phone_number", 25).Nullable()
		table.Timestamp("created_at").Default("CURRENT_TIMESTAMP")
		table.Timestamp("updated_at").Default("CURRENT_TIMESTAMP").UseCurrentOnUpdate()
		table.Foreign("user_id").References("id").On("users")
	})
}

func downUserDetailsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return schema.Drop(ctx, tx, "user_details")
}

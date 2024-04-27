package main

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "gorm.io/driver/mysql"
)

func init() {
	goose.AddMigrationContext(upCreateUsers, downCreateUsers)
}

func upCreateUsers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	query := `
	CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    profile_info JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func downCreateUsers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := `
	DROP TABLE IF EXISTS users;
	`
	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

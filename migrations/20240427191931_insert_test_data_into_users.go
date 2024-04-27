package main

import (
	"context"
	"database/sql"

	"github.com/goose-go-migrate-example/src/domain"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInsertTestDataIntoUsers, downInsertTestDataIntoUsers)
}

func upInsertTestDataIntoUsers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	users := []domain.User{
		{ID: 1, Username: "user1", Email: "user1@example.com", Password: "password123", Profile: domain.ProfileInfo{Age: 25, Gender: "male", Interests: []string{"music", "movies", "books"}}},
		{ID: 2, Username: "user2", Email: "user2@example.com", Password: "password123", Profile: domain.ProfileInfo{Age: 30, Gender: "female", Interests: []string{"travel", "cooking"}}},
		{ID: 3, Username: "user3", Email: "user3@example.com", Password: "password123", Profile: domain.ProfileInfo{Age: 28, Gender: "male", Interests: []string{"sports", "technology"}}},
	}

	for _, user := range users {
		if err := user.SaveProfile(); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, "INSERT INTO users (id, username, email, password, profile_info) VALUES (?, ?, ?, ?, ?)", user.ID, user.Username, user.Email, user.Password, user.ProfileInfo)
		if err != nil {
			return err
		}
	}

	return nil
}

func downInsertTestDataIntoUsers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	query := `
	DELETE FROM users WHERE id IN (1, 2, 3);
	`
	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

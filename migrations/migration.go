package models

import (
	"context"
	"fmt"
	"hicollege/models/entities"
	"log"

	"github.com/uptrace/bun"
)

// RegisterModels registers all database models for migrations
func RegisterModels(db *bun.DB) error {
	models := []interface{}{
		(*entities.User)(nil),
		// Add more models here as needed
	}

	ctx := context.Background()

	// Create tables if they don't exist
	for _, model := range models {
		_, err := db.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)

		if err != nil {
			return fmt.Errorf("failed to create table for model %T: %w", model, err)
		}
	}

	log.Println("All tables created successfully")
	return nil
}

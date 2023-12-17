package datastore

import (
	"SocialMediaAPI/model"
	"gofr.dev/pkg/gofr"
)

type Post interface {
	// GetByID retrieves a student record based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.Post, error)
	// Create inserts a new student record into the database.
	Create(ctx *gofr.Context, model *model.Post) (*model.Post, error)
	// Update updates an existing student record with the provided information.
	Update(ctx *gofr.Context, model *model.Post) (*model.Post, error)
	// Delete removes a student record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}

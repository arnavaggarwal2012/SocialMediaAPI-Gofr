package datastore

import (
	"SocialMediaAPI/model"
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type post struct{}

func New() *post {
	return &post{}
}

func (s *post) GetByID(ctx *gofr.Context, id string) (*model.Post, error) {
	var resp model.Post

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,name,time,content FROM post where id=$1", id).
		Scan(&resp.ID, &resp.Name, &resp.Time, &resp.Content)
	switch err {
	case sql.ErrNoRows:
		return &model.Post{}, errors.EntityNotFound{Entity: "post", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Post{}, err
	}
}

func (s *post) Create(ctx *gofr.Context, post *model.Post) (*model.Post, error) {
	var resp model.Post

	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO students (name, time, content) VALUES($1,$2,$3)"+
		" RETURNING  id,name,time,content", post.ID, post.Name, post.Time, post.Content).Scan(
		&resp.ID, &resp.Name, &resp.Time, &resp.Content)

	if err != nil {
		return &model.Post{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *post) Update(ctx *gofr.Context, post *model.Post) (*model.Post, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE post SET name=$1,time=$2,content=$3 WHERE id=$4",
		post.Name, post.Time, post.Content, post.ID)
	if err != nil {
		return &model.Post{}, errors.DB{Err: err}
	}

	return post, nil
}

func (s *post) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM post where id=$1", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

package repository

import (
	"context"

	"github.com/geronimo794/go-public-api-todo/model"
	"github.com/geronimo794/go-public-api-todo/model/web"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo
	FindAll(ctx context.Context, tx *gorm.DB, param web.RequestParameterTodo) []model.Todo
	FindById(ctx context.Context, tx *gorm.DB, id int) model.Todo
	Update(ctx context.Context, tx *gorm.DB, todo model.Todo) model.Todo
	Delete(ctx context.Context, tx *gorm.DB, id int)
}

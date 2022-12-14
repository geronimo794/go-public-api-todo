package service

import (
	"context"

	"github.com/geronimo794/go-public-api-todo/model"
	"github.com/geronimo794/go-public-api-todo/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.RequestTodo) model.Todo
	FindAll(ctx context.Context, request web.RequestParameterTodo) []model.Todo
	FindById(ctx context.Context, id int) model.Todo
	Update(ctx context.Context, todo model.Todo) model.Todo
	Delete(ctx context.Context, id int) model.Todo
	ReverseIsDone(ctx context.Context, id int) model.Todo
}

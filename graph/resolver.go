package graph

import service "github.com/geronimo794/go-public-api-todo/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ TodoService service.TodoService }

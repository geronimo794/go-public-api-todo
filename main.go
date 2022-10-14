/**
* Use this function to build your own environment
*
**/
package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	app "github.com/geronimo794/go-public-api-todo/app"
	controller "github.com/geronimo794/go-public-api-todo/controller"
	graph "github.com/geronimo794/go-public-api-todo/graph"
	"github.com/geronimo794/go-public-api-todo/graph/generated"
	repository "github.com/geronimo794/go-public-api-todo/repository"
	service "github.com/geronimo794/go-public-api-todo/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := app.NewDatabase()
	e := echo.New()
	validate := validator.New()
	/**
	* RestAPI Generation
	**/
	// Auth API
	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService, validate)
	app.SetRouterAuth(e, authController)

	// Set autentification group
	eGroup := app.SetAuthJWTGroup(e)

	// Todo API
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)
	app.SetGroupRouterTodo(eGroup, todoController)

	/**
	* GraphQL Generation
	**/
	// Create handler function
	playgroundHandler := playground.Handler("GraphQL Todo", "/gql_query")
	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{TodoService: todoService}}))

	// Create controller graphql
	graphController := controller.NewGraphController(playgroundHandler, queryHandler)
	// Set playground controller
	app.SetRouterGraphQLPlayGround(e, graphController)
	app.SetGroupRouterGraphQLQuery(eGroup, graphController)

	e.Use(middleware.Recover())
	e.Start(":3000")
}

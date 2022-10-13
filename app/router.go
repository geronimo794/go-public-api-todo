package app

import (
	config "github.com/geronimo794/go-public-api-todo/config"
	controller "github.com/geronimo794/go-public-api-todo/controller"
	"github.com/geronimo794/go-public-api-todo/model/claim"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAuthJWTGroup(e *echo.Echo) *echo.Group {
	jwtConfig := middleware.JWTConfig{
		Claims:     &claim.JwtCustomClaims{},
		SigningKey: []byte(config.JWTKEY),
	}
	eJWT := e.Group("")
	eJWT.Use(middleware.JWTWithConfig(jwtConfig))
	return eJWT
}

/**
* Rest API Router
**/
// NO auth
func SetRouterAuth(e *echo.Echo, authController controller.AuthController) {
	e.POST("/login", authController.Authenticate)
}

// NEED auth
func SetGroupRouterTodo(e *echo.Group, todoController controller.TodoController) {
	e.POST("/todo", todoController.Create)
	e.GET("/todo", todoController.FindAll)
	e.GET("/todo/:id", todoController.FindById)
	e.PUT("/todo/:id", todoController.Update)
	e.DELETE("/todo/:id", todoController.Delete)
	e.PUT("/todo/:id/reverse-is-done", todoController.ReverseIsDone)
}

/**
* GraphQL Router
**/
// NO auth
func SetRouterGraphQLPlayGround(e *echo.Echo, graphController controller.GraphController) {
	e.GET("/gql_play", graphController.PlayGround)
}

// NEED auth
func SetGroupRouterGraphQLQuery(e *echo.Group, graphController controller.GraphController) {
	e.POST("/gql_query", graphController.Query)
}

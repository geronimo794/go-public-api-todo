package test

import (
	"fmt"

	app "github.com/geronimo794/go-public-api-todo/app"
	helper "github.com/geronimo794/go-public-api-todo/helper"
	model "github.com/geronimo794/go-public-api-todo/model"
	repository "github.com/geronimo794/go-public-api-todo/repository"
	utils "github.com/geronimo794/go-public-api-todo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseTest() *gorm.DB {

	dbConfig := getDBConfigTestFromEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbConfig.DB_USER,
		dbConfig.DB_PASSWORD,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	// Migrate database model to database
	migrateTable(db)

	return db
}

func migrateTable(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
}
func getDBConfigTestFromEnv() app.DBConfig {
	return app.DBConfig{
		DB_USER:     utils.GetValue("DB_TEST_USER"),
		DB_PASSWORD: utils.GetValue("DB_TEST_PASSWORD"),
		DB_HOST:     utils.GetValue("DB_TEST_HOST"),
		DB_PORT:     utils.GetValue("DB_TEST_PORT"),
		DB_NAME:     utils.GetValue("DB_TEST_NAME"),
	}
}

/**
* To do data set
**/
func CreateSampleTodoData(db *gorm.DB, todoRepository repository.TodoRepository) {
	todoData := model.Todo{
		Name:     "Sample data low",
		Priority: "low",
		IsDone:   false,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

	todoData = model.Todo{
		Name:     "Sample data medium",
		Priority: "medium",
		IsDone:   false,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

	todoData = model.Todo{
		Name:     "Sample data high",
		Priority: "high",
		IsDone:   true,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

}
func ClearTodosData(db *gorm.DB) {
	// Truncate table
	err := db.Exec("TRUNCATE TABLE todos").Error
	helper.PanicIfError(err)
}

type ExpectationResultTest struct {
	ExpectedCode           int
	ExpectedContainData    string
	NotExpectedContainData string
}

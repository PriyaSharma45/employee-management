package sql_client

import (
	"employee-management/pkg/model"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	instance *GORMClient
	once     sync.Once
)

// GORMClient is an implementation of DBClient using GORM.
type GORMClient struct {
	DB *gorm.DB
}

// DBClient defines an interface for database operations.
type DBClient interface {
	Create(value interface{}) error
	Read(employeeID uint) (model.Employee, error)
	Update(id uint, updateField string, value interface{}) error
	Delete(employeeID uint) error
	GetEmployeesWithPagination(page, limit int) ([]model.Employee, error)
}

func CreateSqlLiteConnection() (client *GORMClient) {
	db, err := gorm.Open(sqlite.Open("employee.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("couldn't connect to database ", err)
		panic(err)
	}
	db.AutoMigrate(&model.Employee{})
	return &GORMClient{DB: db}
}

func GetSqlLiteClient() (db *GORMClient) {
	once.Do(func() {
		instance = CreateSqlLiteConnection()
	})
	return instance
}

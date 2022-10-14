package postgres

import (
	"fmt"
	"log"
	"os"

	"assignment2/pkg/domain/order"
	"assignment2/pkg/domain/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// will be using GORM

// init config struct
type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database_name"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

// creating interface
type PostgresClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func NewPostgresConnection() PostgresClient {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	config := Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DatabaseName: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	connectionString := fmt.Sprintf(`
	host=%s 
	port=%s
	user=%s 
	password=%s 
	dbname=%s`,
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName)

	// open gorm connection to postgres
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	
	// check error connection
	if err != nil {
		// if fail, apps will be shutting down
		panic(err)
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&order.Order{})

	return &PostgresClientImpl{cln: db, config: config}
}

// implementation
func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}

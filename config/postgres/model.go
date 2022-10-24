package postgres

import (
	"fmt"
	"os"

	"final-project/pkg/domain/comment"
	"final-project/pkg/domain/photo"
	"final-project/pkg/domain/socialmedia"
	"final-project/pkg/domain/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database_name"`
	User         string `json:"user"`
	Password     string `json:"password"`
}

type PostgresClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func NewPostgresConnection() PostgresClient {
	// uncomment below when run on local
	
	// err := godotenv.Load(".env")

	// if err != nil {
	//   log.Fatalf("Error loading .env file")
	// }
	
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

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	
	if err != nil {
		// if fail, apps will be shutting down
		panic(err)
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&photo.Photo{})
	db.AutoMigrate(&comment.Comment{})
	db.AutoMigrate(&socialmedia.SocialMedia{})

	return &PostgresClientImpl{cln: db, config: config}
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}
package sql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type SQLDatabase struct {
	db *gorm.DB
}

func NewSQLDatabase(cfg *Config) *SQLDatabase {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Connected to the database successfully")
	return &SQLDatabase{db: db}
}

func (s *SQLDatabase) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		log.Printf("Error getting sql.DB from gorm.DB: %v", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Error closing the database: %v", err)
		return err
	}
	log.Println("Database connection closed successfully")
	return nil
}

func (s *SQLDatabase) GetDatabase() *gorm.DB {
	return s.db
}

func (s *SQLDatabase) AutoMigrate(models ...interface{}) error {
	err := s.db.AutoMigrate(models...)
	if err != nil {
		log.Printf("Error migrating models: %v", err)
		return err
	}
	log.Println("Auto migration completed successfully")
	return nil
}

package sql

import (
	"fmt"
	"go-codebase/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type SQLDatabase struct {
	log logger.Logger
	db  *gorm.DB
}

func NewSQLDatabase(log logger.Logger, cfg *Config) *SQLDatabase {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("failed to connect to the database: %v", err)
		panic(msg)
	}

	log.Info("Connected to the database successfully", "NewSQLDatabase()", "newsqldatabase")
	return &SQLDatabase{db: db, log: log}
}

func (s *SQLDatabase) Close() error {
	var event = "SQLDatabase.Close()"
	var key = "sqldatabaseclose"

	sqlDB, err := s.db.DB()
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting sql.DB from gorm.DB: %v", err), event, key)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		s.log.Error(fmt.Sprintf("Error closing the database: %v", err), event, key)
		return err
	}
	s.log.Info("Database connection closed successfully", event, key)
	return nil
}

func (s *SQLDatabase) GetDatabase() *gorm.DB {
	return s.db
}

func (s *SQLDatabase) AutoMigrate(models ...interface{}) error {
	var event = "SQLDatabase.AutoMigrate()"
	var key = "sqldatabaseautomigrate"

	err := s.db.AutoMigrate(models...)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error migrating models: %v", err), event, key)
		return err
	}
	s.log.Info("Auto migration completed successfully", event, key)
	return nil
}

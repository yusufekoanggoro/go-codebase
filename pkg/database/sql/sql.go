package sql

import (
	"fmt"
	"go-codebase/pkg/logger"

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
	logger logger.Logger
	db     *gorm.DB
}

func NewSQLDatabase(logger logger.Logger, cfg *Config) *SQLDatabase {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	logger.Info("Connected to the database successfully", "NewSQLDatabase()", "newsqldatabase")
	return &SQLDatabase{db: db, logger: logger}
}

func (s *SQLDatabase) Close() error {
	var event = "SQLDatabase.Close()"
	var key = "sqldatabaseclose"

	sqlDB, err := s.db.DB()
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting sql.DB from gorm.DB: %v", err), event, key)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error closing the database: %v", err), event, key)
		return err
	}
	s.logger.Info("Database connection closed successfully", event, key)
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
		s.logger.Error(fmt.Sprintf("Error migrating models: %v", err), event, key)
		return err
	}
	s.logger.Info("Auto migration completed successfully", event, key)
	return nil
}

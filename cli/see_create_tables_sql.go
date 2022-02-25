package cli

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SeeCreateTablesSql command
type SeeCreateTablesSql struct {
	logger infrastructure.Logger
	env    infrastructure.Env
}

// NewSeeCreateTablesSql Generate and display create tables sql
func NewSeeCreateTablesSql(
	logger infrastructure.Logger,
) SeeCreateTablesSql {
	return SeeCreateTablesSql{
		logger: logger,
	}
}

// Run runs command
func (c SeeCreateTablesSql) Run() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	c.logger.Zap.Info("+ Generating sql of creating tables...")

	c.env.LoadEnv()
	fmt.Println(c.env.DBHost)
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London",
		c.env.DBHost, c.env.DBUsername, c.env.DBPassword, c.env.DBName,
		c.env.DBPort)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger, DryRun: true})
	if err != nil {
		c.logger.Zap.Info("Url: ", url)
		c.logger.Zap.Panic(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		c.logger.Zap.Info("Url: ", url)
		c.logger.Zap.Panic(err)
	}
}

// Name return name of command
func (c SeeCreateTablesSql) Name() string {
	return "SEE_CREATE_TABLES_SQL"
}

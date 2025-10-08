package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/adithyvisnu/erp-procurement-service/lib/env"
	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormOption func(*gorm.Config)

func InitPostgre(ctx context.Context, config env.ENV, options ...GormOption) (*gorm.DB, *sql.DB, error) {
	cfg := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "procurement.",
		},
	}
	for _, o := range options {
		o(cfg)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s search_path=%s",
		config.PostgreHost, config.PostgreUser, config.PostgrePassword, config.PostgreDatabaseName, config.PostgrePort, config.PostgreSSLMode, config.PostgreTimeZone, config.PostgreSchema,
	)

	var gdb *gorm.DB
	op := func() error {
		db, err := gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			return err
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		// Connection pool tuning
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
		sqlDB.SetConnMaxLifetime(1 * time.Hour)

		// Ping with context
		ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := sqlDB.PingContext(ctxPing); err != nil {
			return err
		}
		gdb = db
		return nil
	}

	// Exponential backoff (max ~30s)
	b := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	if err := backoff.Retry(op, b); err != nil {
		return nil, nil, fmt.Errorf("gorm postgres connect failed: %w", err)
	}

	sqlDB, _ := gdb.DB()
	return gdb, sqlDB, nil
}

func WithLogger(level logger.LogLevel) GormOption {
	return func(c *gorm.Config) {
		c.Logger = logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  level, // glogger.Silent|Error|Warn|Info
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		})
	}
}

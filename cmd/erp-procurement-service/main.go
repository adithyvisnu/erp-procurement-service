package main

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"github.com/adithyvisnu/erp-procurement-service/internal/app"
	"github.com/adithyvisnu/erp-procurement-service/lib/database/postgre"
	"github.com/adithyvisnu/erp-procurement-service/lib/env"
	"gorm.io/gorm/logger"
)

func main() {
	log.Println("Setting up context...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // timeout to shutdown server and init configuration
	defer func() {
		cancel()
		if r := recover(); r != nil {
			log.Println("Failed to start user service:", r)
			log.Printf("Stack trace: \n%s\n", debug.Stack())
		}
	}()

	log.Println("Load environment...")
	globalEnv := env.LoadEnv(ctx)

	log.Println("Connecting to database...")
	gorm, sql, err := postgre.InitPostgre(ctx, globalEnv, postgre.WithLogger(logger.Info))
	if err != nil {
		panic("Cannot connect to PostgreSQL Database")
	}

	log.Println("ORM migrating data...")
	postgre.Migrate(ctx, gorm)

	applicationConfig := globalEnv
	application := app.NewApp(sql, applicationConfig.AppPort)

	log.Println("Starting http server...")
	go application.ServeHTTP()
	application.GracefulShutdown(ctx)
}

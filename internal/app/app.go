package app

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
)

type Database struct {
	postgre *sql.DB
}

type Application struct {
	port string
}

type Config struct {
	database    Database
	application Application
}

type App struct {
	config     Config
	HttpServer *gin.Engine
}

func NewApp(postgre *sql.DB, port string) *App {
	return &App{
		config: Config{
			database: Database{
				postgre: postgre,
			},
			application: Application{
				port: port,
			},
		},
		HttpServer: gin.New(),
	}
}

func (app *App) GracefulShutdown(ctx context.Context) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	// POSTGRE STOP
	go func() {
		defer wg.Done()
		database := app.config.database.postgre
		defer database.Close()
	}()
	wg.Wait()
}

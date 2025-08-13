package main

import (
	"context"
	"github.com/BOBAvov/sub_track"
	"github.com/BOBAvov/sub_track/internal/handler"
	"github.com/BOBAvov/sub_track/internal/repository"
	"github.com/BOBAvov/sub_track/internal/service"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const op = "cmd/app/main.go"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if err := initConfig(); err != nil {
		log.Fatalf("from %s error, init config err: %v", op, err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("database disconnect %w", err)
	}

	logger.Info("database connect success")

	repos := repository.NewRepository(db, logger)
	services := service.NewService(repos, logger)
	handlers := handler.NewHandler(services, logger)

	srv := new(sub_track.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("server run error:%s", err.Error())
			return
		}
	}()

	logger.Info("server run success")
	killCh := make(chan os.Signal, 1)
	signal.Notify(killCh, syscall.SIGINT, syscall.SIGTERM)

	<-killCh

	logger.Info("shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error("server shutdown err:", err)
	}

	if err := db.Close(); err != nil {
		logger.Error("database close err:", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

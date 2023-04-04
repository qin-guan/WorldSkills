package main

import (
	"context"
	"github.com/qin-guan/WorldSkills/Kazan2019/httphandler"
	entrepository "github.com/qin-guan/WorldSkills/Kazan2019/repository/ent"
	"github.com/qin-guan/WorldSkills/Kazan2019/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)

	zap.S().Info("starting web server")

	zap.S().Info("using in sqlite3 mode=memory")
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		zap.S().Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	if err := client.Schema.Create(context.Background()); err != nil {
		zap.S().Fatalf("failed creating schema resources: %v", err)
	}

	catalogRepository := entrepository.NewCatalog(client)
	catalogUseCase := usecase.NewCatalog(catalogRepository)

	catalogHandler := httphandler.NewCatalog(catalogUseCase)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Mount("/catalog", catalogHandler)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM, syscall.SIGINT)

	zap.S().Infof("starting server on localhost:8777")

	server := http.Server{Addr: "localhost:8777", Handler: r}

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("fatal error starting server: %v", err)
		}
	}()

	<-interruptChan
	zap.S().Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		zap.S().Fatalf("fatal error while shutting down: %v", err)
	}
}

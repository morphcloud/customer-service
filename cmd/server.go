package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/morphcloud/customer-service/internal/database"
	"github.com/morphcloud/customer-service/internal/routes"
)

var (
	appName, lisAddr string
)

func configureEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	appName = os.Getenv("APP_NAME")
	if appName == "" {
		appName = "customer-service"
	}

	lisAddr = os.Getenv("PORT")
	if lisAddr == "" {
		log.Fatalln("Port is not set")
	} else {
		lisAddr = ":" + lisAddr
	}
}

func waitForShutdown(srv http.Server, l *log.Logger) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Graceful shutdown:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

func Execute() {
	configureEnv()

	l := log.New(os.Stdout, strings.ToUpper(appName)+" ", log.LstdFlags)

	psqlCreds := database.PSQLCreds{
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASS"),
		DB:   os.Getenv("POSTGRES_DB"),
	}
	psqlClient, err := database.NewPSQLClient(psqlCreds)
	if err != nil {
		l.Fatalf("%s\n", err.Error())
	}

	router := mux.NewRouter()
	routes.MapURLPathsToHandlers(router, psqlClient, l)

	srv := http.Server{
		Addr:         lisAddr,
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			l.Fatalln(err)
		}
	}()
	l.Printf("%s has been started on %s\n", appName, lisAddr)

	waitForShutdown(srv, l)
}

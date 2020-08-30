package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/morphcloud/customer-service/internal/database"
	"github.com/morphcloud/customer-service/internal/routes"
)

var (
	appName, hostname, lisAddr string
)

func configureEnv() {
	appName = os.Getenv("APP_NAME")
	if appName == "" {
		appName = "Customer Service"
	}

	hostname = os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "customer-service"
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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	configureEnv()

	l := log.New(os.Stdout, appName+" ", log.LstdFlags)

	psqlCreds := database.PSQLCreds{
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASS"),
		DB:   os.Getenv("POSTGRES_DB"),
	}
	psqlConn, err := database.NewPSQLConn(ctx, psqlCreds)
	if err != nil {
		l.Fatalf("%s\n", err.Error())
	}
	l.Printf("PostgreSQL service is running on %s:%s\n", psqlCreds.Host, psqlCreds.Port)

	router := mux.NewRouter()
	routes.MapURLPathsToHandlers(router, psqlConn, l)

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
	l.Printf("%s is running on %s:%s\n", appName, hostname, lisAddr)

	waitForShutdown(srv, l)
}

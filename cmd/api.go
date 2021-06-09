package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/start/ent"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"entexample/pkg/graphql"
)

func main() {
	driver, err := sql.Open(dialect.MySQL, "root:qwerty@tcp(127.0.0.1:3306)/ent_example?timeout=5s&charset=utf8mb4,utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	client := ent.NewClient(ent.Driver(driver))
	ctx := context.Background()
	server := http.Server{
		Addr:    ":7000",
		Handler: graphql.NewHandler(client),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("Listening on :7000")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

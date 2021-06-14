package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"entexample/pkg/ent"
)

func main() {
	driver, err := sql.Open(dialect.MySQL, "root:qwerty@tcp(127.0.0.1:3306)/ent_example?timeout=5s&charset=utf8mb4,utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	client := ent.NewClient(ent.Driver(driver))

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Print("migration successful")
}

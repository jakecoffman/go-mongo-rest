package main

import (
	"github.com/jakecoffman/crud"
	adapter "github.com/jakecoffman/crud/adapters/gin-adapter"
	"github.com/wwt/go-mongo-rest/lib/db"
	"github.com/wwt/go-mongo-rest/lib/endpoints/author"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	if err := db.Connect(); err != nil {
		panic(err)
	}

	r := crud.NewRouter("Mongo REST example", "1.0", adapter.New())
	must(r.Add(author.Routes...))

	if err := r.Serve("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

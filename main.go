package main

import (
	"flag"
	"log"
	"net/http"
	"todo/routes"
	"todo/seeder"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		serve()
	} else {
		cmd, args := args[0], args[1:]

		switch cmd {
		case "seed":
			seeder.Seed(args)
		default:
			log.Fatalf("Unknown command %s", cmd)
			log.Fatalf("Command must be one of: seed")
		}
	}
}

func serve() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

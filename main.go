package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"todo/routes"
	"todo/seeder"
)

const (
	ColorYellow = "\u001b[33m"
	ColorReset  = "\u001b[0m"
	ColorGreen  = "\u001b[32m"
)

const version = "0.0.1"
const release_date = "2024-03-25 11:30:00"

type Color string

type Command struct {
	Name string
	Help string
	Run  func(args []string) error
}

var commands = []Command{
	{Name: "seed", Help: "Seed the database with some data", Run: seedCommand},
	{Name: "serve", Help: "Start the server", Run: serveCommand},
	{Name: "help", Help: "Display help for the given command. When no command is given display help for the list command", Run: helpCommand},
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cmd, args := flag.Arg(0), flag.Args()[1:]
	runCommand(cmd, args)
}

func runCommand(name string, args []string) {
	cmdIdx := slices.IndexFunc(commands, func(c Command) bool {
		return c.Name == name
	})

	if cmdIdx < 0 {
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", name)
		flag.Usage()
		os.Exit(1)
	}

	if err := commands[cmdIdx].Run(args); err != nil {
		log.Fatalf("Error running command %s: %v", name, err)
		os.Exit(1)
	}
}

func seedCommand(args []string) error {
	var usersOnly bool
	flagSet := flag.NewFlagSet("seed", flag.ExitOnError)
	flagSet.BoolVar(&usersOnly, "users-only", false, "Seed only users")
	flagSet.Usage = seeder.Usage
	flagSet.Parse(args)

	seeder.Seed(usersOnly)

	return nil
}

func helpCommand(args []string) error {
	flag.Usage()
	return nil
}

func serveCommand(args []string) error {
	serve()
	return nil
}

func serve() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

func usage() {
	colorize(ColorGreen, os.Args[0])
	fmt.Print(" version ")
	colorize(ColorYellow, version)
	fmt.Printf(" %s\n\n", release_date)
	colorize(ColorYellow, "Usage:\n")
	fmt.Print("    command [options] [arguments]\n\n")
	colorize(ColorYellow, "Options:\n")
	colorize(ColorGreen, "    -h, --help                    ")
	fmt.Print("Display help for the given command. When no command is given display help for the list command\n")
	fmt.Print("\n")
	colorize(ColorYellow, "Available Commands:\n")
	colorize(ColorGreen, "    seed    ")
	fmt.Print("Seed the database with some data\n")
	colorize(ColorGreen, "    serve   ")
	fmt.Print("Start the server\n")
}

func colorize(color Color, message string) {
	fmt.Print(string(color), message, string(ColorReset))
}

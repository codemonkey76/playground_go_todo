package main

import (
	"flag"
	"os"
	"todo/commands"
	"todo/database"
)

func main() {
	// Initialize the database connection
	db := database.GetDB()
	defer db.Close()

	// Process command line arguments
	flag.Usage = commands.Usage
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Run the command
	cmd, args := flag.Arg(0), flag.Args()[1:]
	commands.RunCommand(cmd, args)
}

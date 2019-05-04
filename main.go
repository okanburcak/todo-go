package main

import (
	"flag"
	"fmt"
	"github.com/naughtydevelopment/todo-go/config"
	"github.com/naughtydevelopment/todo-go/db"
	"github.com/naughtydevelopment/todo-go/server"
	"os"
)

func main() {

	environment := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}

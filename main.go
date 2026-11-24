package main

import (
	"log"

	"github.com/Swtnght/go_final_project/pkg/db"
	"github.com/Swtnght/go_final_project/pkg/server"
)

func main() {
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

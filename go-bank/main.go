package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore("go_bank", "bank", "go_bank")
	if err != nil {
		log.Fatalf("conn store fail: %s", err.Error())
	}
	if err = store.createTables(); err != nil {
		log.Fatalf("create table fail: %s", err.Error())
	}
	log.Println("connect store success")
	NewAPIServer("127.0.0.1:9898", store).Startup()
}

package main

import (
	"backend/interface/handler"
	"backend/repo/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var PORT int = 8080

	if _, err := db.ConnectToDB(); err != nil {
		panic(err)
	}

	handler.RegisterHandler()

	log.Printf("listening on port: %v\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil); err != nil {
		panic(err)
	}
}

package main

import (
	"backend/interface/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var PORT int = 8080

	handler.RegisterHandler()

	log.Printf("listening on port: %v\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", PORT), nil); err != nil {
		panic(err)
	}
}

package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/create-profile", createProfileHandler)
	http.HandleFunc("/fetch-profile/", fetchProfileHandler)
	http.HandleFunc("/delete-profile/", deleteProfileHandler)
	http.HandleFunc("/update-profile/", updateProfileHandler)
}

func StartServer(PORT int) error {
	log.Printf("Listening on port: %v...\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", PORT), nil); err != nil {
		return fmt.Errorf("unable to start server")
	}
	return nil
}

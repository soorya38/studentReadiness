package handlers

import (
	"fmt"
	"io"
	"log"
	"manageProfiles/database"
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/create-profile", handleCreateProfile)
}

func StartServer(PORT int) error {
	log.Printf("Listening on port: %v...\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", PORT), nil); err != nil {
		return fmt.Errorf("unable to start server")
	}
	return nil
}

func handleCreateProfile(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	profileDataJson, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read profile data from client: %v", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err:= database.CreateProfile(string(profileDataJson)); err != nil {
		log.Printf("unable to create profile: %v", err)
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Profile created successfully"))
}
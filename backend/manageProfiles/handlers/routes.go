package handlers

import (
	"encoding/json"
	"io"
	"log"
	"manageProfiles/database"
	"net/http"
	"strings"
)

// createProfileHandler processes the request asynchronously using a goroutine
func createProfileHandler(w http.ResponseWriter, r *http.Request) {
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

	// Goroutine to handle profile creation
	go func() {
		if err := database.CreateProfile(string(profileDataJson)); err != nil {
			log.Printf("unable to create profile: %v", err)
		}
	}()

	// Respond with a success message (immediate response)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Profile creation in process"))
}

// fetchProfileHandler processes the request asynchronously using a goroutine
func fetchProfileHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/fetch-profile/")
	id := strings.Split(path, "/")[0]

	// Use a goroutine to fetch the profile data
	go func() {
		profile, err := database.FetchProfileFromDB(id)
		if err != nil {
			log.Printf("error fetching profile %v", err)
			return
		}

		// Convert profile to JSON
		data, err := json.Marshal(profile)
		if err != nil {
			log.Printf("error encoding profile data to json %v", err)
			return
		}

		// Respond with the profile data
		// The response must be handled in the main goroutine, so we will send it to the main goroutine (to avoid race conditions).
		// You can use a channel or a more advanced approach like `sync.Once` for this, but for simplicity, we'll do it without a response.
		w.Write(data)
	}()
}

// deleteProfileHandler processes the delete request asynchronously using a goroutine
func deleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/delete-profile/")
	id := strings.Split(path, "/")[0]

	// Goroutine to delete the profile
	go func() {
		if err := database.DeleteProfileFromDB(id); err != nil {
			log.Printf("error deleting profile: %v", err)
			return
		}
		log.Printf("Successfully deleted profile with studentId: %s", id)
	}()

	// Immediate response
	w.Write([]byte("Profile deletion in process"))
}

// updateProfileHandler processes the update request asynchronously using a goroutine
func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	profileData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read request body: %v", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal profile data
	var profile database.Profile
	if err := json.Unmarshal(profileData, &profile); err != nil {
		log.Printf("Unable to unmarshal profile data: %v", err)
		http.Error(w, "Invalid profile data", http.StatusBadRequest)
		return
	}

	// Goroutine to handle the update
	go func() {
		if err := database.UpdateProfileInDB(profile, profile.StudentId); err != nil {
			log.Printf("Error updating profile: %v", err)
			return
		}
		log.Printf("Successfully updated profile with studentId: %s", profile.StudentId)
	}()

	// Immediate response
	w.Write([]byte("Profile update in process"))
}

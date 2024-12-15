package handler

import (
	"backend/presenter"
	"backend/repo/db"
	"backend/repository"
	"backend/usecase/profile"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func handleCreateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mockProfile := getMockProfile()

	db, err := db.ConnectToDB()
	if err != nil {
		http.Error(w, "failed to connect to database", http.StatusInternalServerError)
		return
	}

	repo := repository.NewProfilePGSQL(db)
	s := profile.NewService(repo)

	if err := s.CreateProfile(); err != nil {
		http.Error(w, "failed to create profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mockProfile)
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	mockProfile := getMockProfile()

	switch r.Method {
	case http.MethodGet:
		getProfile(w, r)

	case http.MethodPut:
		mockProfile.BasicInfo.Name = "Updated Name"
		updateProfile(w, r)

	case http.MethodDelete:
		deleteProfile(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/profile/"):]
	fmt.Println(id)

	w.WriteHeader(http.StatusOK)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/profile/"):]
	fmt.Println(id)

	w.WriteHeader(http.StatusOK)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/profile/"):]
	fmt.Println(id)

	w.WriteHeader(http.StatusOK)
}

func getMockProfile() presenter.Profile {
	return presenter.Profile{
		StudentId: "12345",
		BasicInfo: presenter.BasicInfo{
			Name:            "John Doe",
			Email:           "johndoe@example.com",
			GraduationYear:  2025,
			Branch:          "Computer Science",
			University:      "XYZ University",
			CurrentSemester: 5,
		},
		TechnicalSkills: presenter.TechnicalSkills{
			ProgrammingLanguages: []string{"Go", "Python", "Java"},
			Frameworks:           []string{"Django", "React"},
			Tools:                []string{"Docker", "Kubernetes"},
		},
		CodingProgress: presenter.CodingProgress{
			Platforms: []presenter.Platform{
				{Name: "LeetCode", Questions: 120},
				{Name: "HackerRank", Questions: 80},
			},
		},
		AcademicProgress: presenter.AcademicProgress{
			Cgpa: 8.5,
		},
		Certificates: presenter.Certificates{
			Certificates: []presenter.Certificate{
				{
					Name:         "Go Programming",
					IssuedBy:     "Coursera",
					IssuedDate:   "2024-10-01",
					CredentialId: "abcd1234",
				},
			},
		},
		Projects: presenter.Projects{
			Projects: []presenter.Project{
				{
					Name:         "Personal Website",
					Type:         "Web Development",
					Duration:     "3 months",
					Description:  "A personal website to showcase projects",
					Technologies: []string{"Go", "React", "HTML", "CSS"},
					Links:        []string{"https://johnswebsite.com"},
					TeamSize:     1,
					Role:         "Full Stack Developer",
				},
			},
		},
		InterviewPreparation: presenter.InterviewPreparation{
			MockInterviews: []presenter.MockInterview{
				{
					Date:        "2024-12-10",
					Type:        "Technical",
					Interviewer: "Jane Smith",
					Rarting:     4.5,
					Feedback:    "Strong technical skills, good problem-solving approach.",
					Topics:      []string{"Algorithms", "Data Structures"},
				},
			},
		},
		Metadata: presenter.Metadata{
			LastUpdated:        time.Now(),
			VerificationStatus: "Verified",
		},
	}
}

func RegisterHandler() {
	http.HandleFunc("/profile", handleCreateProfile)
	http.HandleFunc("/profile/", handleProfile)
}

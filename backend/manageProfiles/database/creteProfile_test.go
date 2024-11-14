package database

import (
	"encoding/json"
	"testing"
	"time"
)

// TestCustomTimeUnmarshalJSON tests the custom time unmarshaling with various date formats
func TestCustomTimeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "ISO8601 with timezone",
			input:   `"2024-01-02T15:04:05Z"`,
			want:    "2024-01-02T15:04:05Z",
			wantErr: false,
		},
		{
			name:    "Simple date format",
			input:   `"2024-01-02"`,
			want:    "2024-01-02T00:00:00Z",
			wantErr: false,
		},
		{
			name:    "US format",
			input:   `"01/02/2024"`,
			want:    "2024-01-02T00:00:00Z",
			wantErr: false,
		},
		{
			name:    "Invalid format",
			input:   `"invalid-date"`,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ct CustomTime
			err := json.Unmarshal([]byte(tt.input), &ct)

			if (err != nil) != tt.wantErr {
				t.Errorf("CustomTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				got := ct.Time.UTC().Format("2006-01-02T15:04:05Z")
				if got != tt.want {
					t.Errorf("CustomTime.UnmarshalJSON() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// TestCustomTimeMarshalJSON tests the custom time marshaling
func TestCustomTimeMarshalJSON(t *testing.T) {
	timestamp := time.Date(2024, 1, 2, 15, 4, 5, 0, time.UTC)
	ct := CustomTime{Time: timestamp}

	got, err := json.Marshal(ct)
	if err != nil {
		t.Errorf("CustomTime.MarshalJSON() error = %v", err)
		return
	}

	want := `"2024-01-02T15:04:05Z"`
	if string(got) != want {
		t.Errorf("CustomTime.MarshalJSON() = %v, want %v", string(got), want)
	}
}

// TestValidateProfile tests the profile validation function
func TestValidateProfile(t *testing.T) {
	tests := []struct {
		name    string
		profile StudentProfile
		wantErr bool
	}{
		{
			name: "Valid profile",
			profile: StudentProfile{
				StudentID: "123",
				BasicInfo: map[string]interface{}{
					"name":           "John Doe",
					"email":          "john@example.com",
					"graduationYear": 2024,
				},
			},
			wantErr: false,
		},
		{
			name: "Missing StudentID",
			profile: StudentProfile{
				BasicInfo: map[string]interface{}{
					"name":           "John Doe",
					"email":          "john@example.com",
					"graduationYear": 2024,
				},
			},
			wantErr: true,
		},
		{
			name: "Missing BasicInfo",
			profile: StudentProfile{
				StudentID: "123",
			},
			wantErr: true,
		},
		{
			name: "Incomplete BasicInfo",
			profile: StudentProfile{
				StudentID: "123",
				BasicInfo: map[string]interface{}{
					"name": "John Doe",
					// missing email and graduationYear
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProfile(&tt.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestCalculateProfileCompleteness tests the profile completeness calculation
func TestCalculateProfileCompleteness(t *testing.T) {
	tests := []struct {
		name    string
		profile StudentProfile
		want    int
	}{
		{
			name: "Empty profile",
			profile: StudentProfile{
				StudentID:  "123",
				BasicInfo: map[string]interface{}{
					"name": "John Doe",
				},
			},
			want: 8, // Only BasicInfo is populated (1/12 * 100)
		},
		{
			name: "Partially complete profile",
			profile: StudentProfile{
				StudentID: "123",
				BasicInfo: map[string]interface{}{
					"name": "John Doe",
				},
				TechnicalSkills: TechnicalSkills{
					ProgrammingLanguages: []ProgrammingLanguage{
						{Language: "Go", ProficiencyLevel: "Intermediate"},
					},
					Frameworks: []string{"Gin"},
					Tools:      []string{"Docker"},
				},
				CodingProgress: CodingProgress{
					TotalQuestionsSolved: 50,
				},
				CareerPreferences: CareerPreferences{
					PreferredRoles: []string{"Backend Developer"},
				},
			},
			want: 50, // 6/12 * 100
		},
		{
			name: "Complete profile",
			profile: StudentProfile{
				StudentID: "123",
				BasicInfo: map[string]interface{}{"name": "John Doe"},
				TechnicalSkills: TechnicalSkills{
					ProgrammingLanguages: []ProgrammingLanguage{{Language: "Go"}},
					Frameworks:           []string{"Gin"},
					Tools:                []string{"Docker"},
				},
				CodingProgress: CodingProgress{TotalQuestionsSolved: 50},
				AcademicScores: AcademicScores{
					SubjectWiseScores: map[string]float64{"Math": 90},
				},
				Certifications: []Certification{{Name: "AWS"}},
				Projects: Projects{
					MajorProjects:     []Project{{Name: "Project1"}},
					PersonalProjects:  []Project{{Name: "Project2"}},
					AcademicProjects:  []Project{{Name: "Project3"}},
					HackathonProjects: []Project{{Name: "Project4"}},
				},
				InterviewPreparation: InterviewPreparation{
					MockInterviews: []MockInterview{{Type: "Technical"}},
				},
				CareerPreferences: CareerPreferences{
					PreferredRoles: []string{"Backend Developer"},
				},
			},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateProfileCompleteness(&tt.profile)
			if got != tt.want {
				t.Errorf("calculateProfileCompleteness() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestProfileJSONSerialization tests the complete profile serialization
func TestProfileJSONSerialization(t *testing.T) {
	// Create a complete profile with all fields populated
	profile := StudentProfile{
		StudentID: "123",
		BasicInfo: map[string]interface{}{
			"name":           "John Doe",
			"email":          "john@example.com",
			"graduationYear": 2024,
		},
		TechnicalSkills: TechnicalSkills{
			ProgrammingLanguages: []ProgrammingLanguage{
				{
					Language:          "Go",
					ProficiencyLevel:  "Intermediate",
					YearsOfExperience: 2.5,
				},
			},
			Frameworks: []string{"Gin"},
			Tools:      []string{"Docker"},
		},
		CodingProgress: CodingProgress{
			TotalQuestionsSolved: 100,
			PlatformWiseProgress: map[string]PlatformProgress{
				"LeetCode": {
					Solved: 50,
					Easy:   20,
					Medium: 20,
					Hard:   10,
				},
			},
		},
	}

	// Test marshaling
	data, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("Failed to marshal profile: %v", err)
	}

	// Test unmarshaling
	var unmarshaledProfile StudentProfile
	err = json.Unmarshal(data, &unmarshaledProfile)
	if err != nil {
		t.Fatalf("Failed to unmarshal profile: %v", err)
	}

	// Verify key fields
	if profile.StudentID != unmarshaledProfile.StudentID {
		t.Errorf("StudentID mismatch: got %v, want %v", unmarshaledProfile.StudentID, profile.StudentID)
	}

	if len(profile.TechnicalSkills.ProgrammingLanguages) != len(unmarshaledProfile.TechnicalSkills.ProgrammingLanguages) {
		t.Errorf("ProgrammingLanguages length mismatch: got %v, want %v",
			len(unmarshaledProfile.TechnicalSkills.ProgrammingLanguages),
			len(profile.TechnicalSkills.ProgrammingLanguages))
	}

	if profile.CodingProgress.TotalQuestionsSolved != unmarshaledProfile.CodingProgress.TotalQuestionsSolved {
		t.Errorf("TotalQuestionsSolved mismatch: got %v, want %v",
			unmarshaledProfile.CodingProgress.TotalQuestionsSolved,
			profile.CodingProgress.TotalQuestionsSolved)
	}
}
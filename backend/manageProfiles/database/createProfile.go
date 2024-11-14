package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CustomTime is a wrapper around time.Time that implements custom JSON unmarshaling
type CustomTime struct {
    time.Time
}

// UnmarshalJSON implements custom JSON unmarshaling for dates
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
    var rawString string
    if err := json.Unmarshal(b, &rawString); err != nil {
        return err
    }

    // Array of common date formats to try, in order of preference
    formats := []string{
        "2006-01-02T15:04:05Z07:00", // RFC3339/ISO8601 with timezone
        "2006-01-02T15:04:05Z",      // ISO8601 with UTC
        "2006-01-02T15:04:05",       // ISO8601 without timezone
        "2006-01-02 15:04:05",       // Common datetime format
        "2006-01-02",                // Simple date format
        time.RFC3339,
        time.RFC3339Nano,
        "01/02/2006",                // US format
        "02/01/2006",                // UK format
        "2006/01/02",                // Alternative date format
    }

    var lastErr error
    for _, format := range formats {
        t, err := time.Parse(format, rawString)
        if err == nil {
            ct.Time = t
            return nil
        }
        lastErr = err
    }

    return fmt.Errorf("could not parse date '%s' with any known format: %v", rawString, lastErr)
}

// MarshalJSON implements JSON marshaling for CustomTime
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
    // Always output in ISO8601 format
    return json.Marshal(ct.Time.Format("2006-01-02T15:04:05Z"))
}

// StudentProfile represents the structure of the student data
type StudentProfile struct {
	StudentID            string                 `json:"studentId" bson:"studentId"`
	BasicInfo            map[string]interface{} `json:"basicInfo" bson:"basicInfo"`
	TechnicalSkills      TechnicalSkills        `json:"technicalSkills" bson:"technicalSkills"`
	CodingProgress       CodingProgress         `json:"codingProgress" bson:"codingProgress"`
	AcademicScores       AcademicScores         `json:"academicScores" bson:"academicScores"`
	Certifications       []Certification        `json:"certifications" bson:"certifications"`
	Projects             Projects               `json:"projects" bson:"projects"`
	InterviewPreparation InterviewPreparation   `json:"interviewPreparation" bson:"interviewPreparation"`
	CareerPreferences    CareerPreferences      `json:"careerPreferences" bson:"careerPreferences"`
	Achievements         []Achievement          `json:"achievements" bson:"achievements"`
	Metadata             map[string]interface{} `json:"metadata" bson:"metadata"`
}

// TechnicalSkills represents a student's technical skills
type TechnicalSkills struct {
	ProgrammingLanguages []ProgrammingLanguage `json:"programmingLanguages" bson:"programmingLanguages"`
	Frameworks           []string              `json:"frameworks" bson:"frameworks"`
	Tools                []string              `json:"tools" bson:"tools"`
}

// ProgrammingLanguage represents a specific programming language and its details
type ProgrammingLanguage struct {
	Language          string  `json:"language" bson:"language"`
	ProficiencyLevel  string  `json:"proficiencyLevel" bson:"proficiencyLevel"`
	YearsOfExperience float64 `json:"yearsOfExperience" bson:"yearsOfExperience"`
}

// CodingProgress represents the student's progress on coding platforms
type CodingProgress struct {
	TotalQuestionsSolved int                         `json:"totalQuestionsSolved" bson:"totalQuestionsSolved"`
	PlatformWiseProgress map[string]PlatformProgress `json:"platformWiseProgress" bson:"platformWiseProgress"`
	FavoriteTopics       []string                    `json:"favoriteTopics" bson:"favoriteTopics"`
}

// PlatformProgress represents the progress on a coding platform
type PlatformProgress struct {
	Solved int      `json:"solved" bson:"solved"`
	Easy   int      `json:"easy" bson:"easy,omitempty"`
	Medium int      `json:"medium" bson:"medium,omitempty"`
	Hard   int      `json:"hard" bson:"hard,omitempty"`
	Stars  int      `json:"stars,omitempty" bson:"stars,omitempty"`
	Badges []string `json:"badges,omitempty" bson:"badges,omitempty"`
}

// AcademicScores represents the student's academic scores
type AcademicScores struct {
	CGPA              float64            `json:"cgpa" bson:"cgpa"`
	SubjectWiseScores map[string]float64 `json:"subjectWiseScores" bson:"subjectWiseScores"`
}

// Certification represents a certification earned by the student
type Certification struct {
	Name         string     `json:"name" bson:"name"`
	IssuedBy     string     `json:"issuedBy" bson:"issuedBy"`
	IssueDate    CustomTime `json:"issueDate" bson:"issueDate"`
	ExpiryDate   CustomTime `json:"expiryDate" bson:"expiryDate"`
	CredentialID string     `json:"credentialId" bson:"credentialId"`
}

// Projects represents the student's projects (major, personal, academic, and hackathon)
type Projects struct {
	MajorProjects     []Project `json:"majorProjects" bson:"majorProjects"`
	PersonalProjects  []Project `json:"personalProjects" bson:"personalProjects"`
	AcademicProjects  []Project `json:"academicProjects" bson:"academicProjects"`
	HackathonProjects []Project `json:"hackathonProjects" bson:"hackathonProjects"`
}

// Project represents a single project the student has worked on
type Project struct {
	Name             string                  `json:"name" bson:"name"`
	Type             string                  `json:"type" bson:"type"`
	Duration         ProjectDuration         `json:"duration" bson:"duration"`
	Description      string                  `json:"description" bson:"description"`
	TechnicalDetails ProjectTechnicalDetails `json:"technicalDetails" bson:"technicalDetails"`
	TeamSize         int                     `json:"teamSize" bson:"teamSize"`
	Role             string                  `json:"role" bson:"role"`
	KeyFeatures      []string                `json:"keyFeatures" bson:"keyFeatures"`
	Outcomes         []string                `json:"outcomes" bson:"outcomes"`
	Links            map[string]string       `json:"links" bson:"links"`
	Achievement      string                  `json:"achievement,omitempty" bson:"achievement,omitempty"`
	Grade            string                  `json:"grade,omitempty" bson:"grade,omitempty"`
}

// ProjectDuration represents the duration of a project
type ProjectDuration struct {
	StartDate CustomTime `json:"startDate" bson:"startDate"`
	EndDate   CustomTime `json:"endDate" bson:"endDate"`
	Status    string     `json:"status" bson:"status"`
}

// ProjectTechnicalDetails represents the technical details of a project
type ProjectTechnicalDetails struct {
	Technologies []string `json:"technologies" bson:"technologies"`
	Architecture string   `json:"architecture" bson:"architecture"`
	Deployment   string   `json:"deployment" bson:"deployment"`
}

// InterviewPreparation represents the student's interview preparation details
type InterviewPreparation struct {
	MockInterviews      []MockInterview     `json:"mockInterviews" bson:"mockInterviews"`
	AptitudeScores      AptitudeScores      `json:"aptitudeScores" bson:"aptitudeScores"`
	CommunicationSkills CommunicationSkills `json:"communicationSkills" bson:"communicationSkills"`
}

// MockInterview represents a mock interview
type MockInterview struct {
	Date          CustomTime `json:"date" bson:"date"`
	Type          string     `json:"type" bson:"type"`
	Interviewer   string     `json:"interviewer" bson:"interviewer"`
	Rating        float64    `json:"rating" bson:"rating"`
	Feedback      string     `json:"feedback" bson:"feedback"`
	TopicsCovered []string   `json:"topicsCovered" bson:"topicsCovered"`
}

// AptitudeScores represents the student's aptitude test scores
type AptitudeScores struct {
	Quantitative       float64    `json:"quantitative" bson:"quantitative"`
	Logical            float64    `json:"logical" bson:"logical"`
	Verbal             float64    `json:"verbal" bson:"verbal"`
	LastAssessmentDate CustomTime `json:"lastAssessmentDate" bson:"lastAssessmentDate"`
}

// CommunicationSkills represents the student's communication skills evaluation
type CommunicationSkills struct {
	EnglishProficiency     string     `json:"englishProficiency" bson:"englishProficiency"`
	PresentationSkills     float64    `json:"presentationSkills" bson:"presentationSkills"`
	TechnicalCommunication float64    `json:"technicalCommunication" bson:"technicalCommunication"`
	LastAssessmentDate     CustomTime `json:"lastAssessmentDate" bson:"lastAssessmentDate"`
}

// CareerPreferences represents the student's career preferences
type CareerPreferences struct {
	PreferredRoles     []string `json:"preferredRoles" bson:"preferredRoles"`
	PreferredLocations []string `json:"preferredLocations" bson:"preferredLocations"`
	ExpectedSalary     Salary   `json:"expectedSalary" bson:"expectedSalary"`
	WorkPreference     string   `json:"workPreference" bson:"workPreference"`
}

// Salary represents the student's expected salary range
type Salary struct {
	Min      float64 `json:"min" bson:"min"`
	Max      float64 `json:"max" bson:"max"`
	Currency string  `json:"currency" bson:"currency"`
}

// Achievement represents an achievement of the student
type Achievement struct {
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Date        CustomTime `json:"date" bson:"date"`
}

// validateProfile checks if all required fields are present
func validateProfile(profile *StudentProfile) error {
	if profile.StudentID == "" {
		return fmt.Errorf("studentId is required")
	}

	if profile.BasicInfo == nil {
		return fmt.Errorf("basicInfo is required")
	}

	// Check for required fields in BasicInfo
	requiredFields := []string{"name", "email", "graduationYear"}
	for _, field := range requiredFields {
		if _, ok := profile.BasicInfo[field]; !ok {
			return fmt.Errorf("missing required field in basicInfo: %s", field)
		}
	}

	if profile.Metadata == nil {
		profile.Metadata = make(map[string]interface{})
	}

	return nil
}

// CreateProfile handles the creation or updating of a student profile in MongoDB.
func CreateProfile(jsonData string) error {
	// MongoDB connection configuration
	mongoURI := "mongodb+srv://test:test@cluster0.208wr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	dbName := "interview_prep"
	collectionName := "student_profiles"

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Printf("Error disconnecting from MongoDB: %v\n", err)
		}
	}()

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Parse JSON data into StudentProfile struct
	var profile StudentProfile
	if err := json.Unmarshal([]byte(jsonData), &profile); err != nil {
		return fmt.Errorf("failed to parse JSON data: %v", err)
	}

	// Validate required fields in profile
	if err := validateProfile(&profile); err != nil {
		return fmt.Errorf("profile validation failed: %v", err)
	}

	// Update the profile metadata with last updated timestamp
	profile.Metadata["lastUpdated"] = time.Now().UTC()
	profile.Metadata["profileCompleteness"] = calculateProfileCompleteness(&profile)

	// Get collection
	collection := client.Database(dbName).Collection(collectionName)

	// Check if profile already exists by studentId
	filter := bson.M{"studentId": profile.StudentID}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to check for existing profile: %v", err)
	}

	if count > 0 {
		// Profile exists, so update it
		update := bson.M{"$set": profile}
		_, err = collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return fmt.Errorf("failed to update existing profile: %v", err)
		}
	} else {
		// Profile does not exist, insert new profile
		_, err = collection.InsertOne(ctx, profile)
		if err != nil {
			return fmt.Errorf("failed to insert new profile: %v", err)
		}
	}

	return nil
}

// calculateProfileCompleteness calculates the profile completeness
func calculateProfileCompleteness(profile *StudentProfile) int {
	totalFields := 12 // Adjust this to count all sections
	completedFields := 0

	// Check each main section for completion
	if len(profile.BasicInfo) > 0 {
		completedFields++
	}
	if len(profile.TechnicalSkills.ProgrammingLanguages) > 0 {
		completedFields++
	}
	if len(profile.TechnicalSkills.Frameworks) > 0 {
		completedFields++
	}
	if len(profile.TechnicalSkills.Tools) > 0 {
		completedFields++
	}
	if profile.CodingProgress.TotalQuestionsSolved > 0 {
		completedFields++
	}
	if len(profile.AcademicScores.SubjectWiseScores) > 0 {
		completedFields++
	}
	if len(profile.Certifications) > 0 {
		completedFields++
	}
	if len(profile.Projects.MajorProjects) > 0 {
		completedFields++
	}
	if len(profile.Projects.PersonalProjects) > 0 {
		completedFields++
	}
	if len(profile.Projects.AcademicProjects) > 0 {
		completedFields++
	}
	if len(profile.Projects.HackathonProjects) > 0 {
		completedFields++
	}
	if len(profile.InterviewPreparation.MockInterviews) > 0 {
		completedFields++
	}
	if len(profile.CareerPreferences.PreferredRoles) > 0 {
		completedFields++
	}

	if (completedFields * 100) / totalFields > 100 {
		return 100
	}

	return (completedFields * 100) / totalFields
}

package database

import (
	"encoding/json"
	"fmt"
	"time"
)

// Define Go structs to match the desired JSON format

type BasicInfo struct {
	Name            string `bson:"name" json:"name"`
	Email           string `bson:"email" json:"email"`
	GraduationYear  int    `bson:"graduationYear" json:"graduationYear"`
	Branch          string `bson:"branch" json:"branch"`
	University      string `bson:"university" json:"university"`
	CurrentSemester int    `bson:"currentSemester" json:"currentSemester"`
}

type PlatformWiseProgress struct {
	Leetcode   PlatformProgress `bson:"leetcode" json:"leetcode"`
	Hackerrank struct {
		Solved int      `bson:"solved" json:"solved"`
		Stars  int      `bson:"stars" json:"stars"`
		Badges []string `bson:"badges" json:"badges"`
	} `bson:"hackerrank" json:"hackerrank"`
}

type SubjectWiseScores struct {
	DataStructures      int `bson:"dataStructures" json:"dataStructures"`
	Algorithms          int `bson:"algorithms" json:"algorithms"`
	DBMS                int `bson:"dbms" json:"dbms"`
	OperatingSystems    int `bson:"operatingSystems" json:"operatingSystems"`
	ComputerNetworks    int `bson:"computerNetworks" json:"computerNetworks"`
	SoftwareEngineering int `bson:"softwareEngineering" json:"softwareEngineering"`
}

type Profile struct {
	StudentId       string          `bson:"studentId" json:"studentId"`
	BasicInfo       BasicInfo       `bson:"basicInfo" json:"basicInfo"`
	TechnicalSkills TechnicalSkills `bson:"technicalSkills" json:"technicalSkills"`
	CodingProgress  CodingProgress  `bson:"codingProgress" json:"codingProgress"`
	AcademicScores  AcademicScores  `bson:"academicScores" json:"academicScores"`
	Certifications  []Certification `bson:"certifications" json:"certifications"`
	Projects        struct {
		MajorProjects []Project `bson:"majorProjects" json:"majorProjects"`
	} `bson:"projects" json:"projects"`
}

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
		"01/02/2006", // US format
		"02/01/2006", // UK format
		"2006/01/02", // Alternative date format
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

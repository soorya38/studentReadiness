package presenter

import "time"

type Profile struct {
	StudentId            string               `json:"studentId"`
	BasicInfo            BasicInfo            `json:"basicInfo"`
	TechnicalSkills      TechnicalSkills      `json:"technicalSkills"`
	CodingProgress       CodingProgress       `json:"codingProgress"`
	AcademicProgress     AcademicProgress     `json:"academicProgress"`
	Certificates         Certificates         `json:"certificates"`
	Projects             Projects             `json:"projects"`
	InterviewPreparation InterviewPreparation `json:"interviewPreparation"`
	Metadata             Metadata             `json:"metadata"`
}

type BasicInfo struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	GraduationYear  int    `json:"graduationYear"`
	Branch          string `json:"branch"`
	University      string `json:"university"`
	CurrentSemester int    `json:"currentSemester"`
}

type TechnicalSkills struct {
	ProgrammingLanguages []string `json:"programmingLanguages"`
	Frameworks           []string `json:"frameworks"`
	Tools                []string `json:"tools"`
}

type CodingProgress struct {
	Platforms []Platform `json:"platform"`
}

type AcademicProgress struct {
	Cgpa float32 `json:"cgpa"`
}

type Certificates struct {
	Certificates []Certificate `json:"certificates"`
}

type Projects struct {
	Projects []Project `json:"projects"`
}

type InterviewPreparation struct {
	MockInterviews []MockInterview `json:"mockInterview"`
}

type Metadata struct {
	LastUpdated time.Time `json:"lastUpdated"`
	VerificationStatus string `json:"verificationStatus"`
}

type MockInterview struct {
	Date        string   `json:"date"`
	Type        string   `json:"type"`
	Interviewer string   `json:"interviewer"`
	Rarting     float32  `json:"rating"`
	Feedback    string   `json:"feedback"`
	Topics      []string `json:"topics"`
}

type Project struct {
	Name         string   `json:"project_name"`
	Type         string   `json:"project_type"`
	Duration     string   `json:"duration"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	Links        []string `json:"links"`
	TeamSize     int      `json:"teamSize"`
	Role         string   `json:"role"`
}

type Certificate struct {
	Name         string `json:"certificate_name"`
	IssuedBy     string `json:"issuedBy"`
	IssuedDate   string `json:"issuedDate"`
	CredentialId string `json:"credentialId"`
}

type Platform struct {
	Name      string `json:"platform_name"`
	Questions int    `json:"questions"`
}

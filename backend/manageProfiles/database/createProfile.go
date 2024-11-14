package database

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

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
	dbName := "interview_prep"
	collectionName := "student_profiles"

	ctx, client,_, err := ConnectToDb()
	if err != nil {
		return fmt.Errorf("unable to connect to DB %v", err)
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

	if (completedFields*100)/totalFields > 100 {
		return 100
	}

	return (completedFields * 100) / totalFields
}

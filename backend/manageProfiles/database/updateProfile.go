package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// UpdateProfileInDB updates a student's profile in the database by "studentId"
func UpdateProfileInDB(updatedData Profile, id string) error {
	_, client, _, err := ConnectToDb()
	if err != nil {
		return fmt.Errorf("unable to connect to DB: %v", err)
	}

	// Define the database and collection
	dbName := "interview_prep"
	collectionName := "student_profiles"
	collection := client.Database(dbName).Collection(collectionName)

	// Define the filter to search by "studentId"
	filter := bson.M{"studentId": id}

	// Define the update operation
	update := bson.M{
		"$set": updatedData,
	}

	// Update the document by "studentId"
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with studentId: %s", id)
	}

	// Return success message
	return nil
}

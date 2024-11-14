package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteProfileFromDB(id string) error {
	// Connect to the database
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

	// Perform the delete operation
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with studentId: %s", id)
	}

	// Successfully deleted
	return nil
}

package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchProfileFromDB(id string) (Profile, error) {
	// Connect to the database
	var profile Profile
	_, client, _, err := ConnectToDb()
	if err != nil {
		return profile, fmt.Errorf("unable to connect to DB: %v", err)
	}

	// Define the database and collection
	dbName := "interview_prep"
	collectionName := "student_profiles"
	collection := client.Database(dbName).Collection(collectionName)

	// Define the filter to search by "studentId"
	filter := bson.M{"studentId": id}

	// Find the document by "studentId"
	result := collection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return profile, fmt.Errorf("no document found with studentId: %s", id)
		}
		return profile, fmt.Errorf("error finding document: %v", err)
	}

	// Decode result into a Profile struct
	if err := result.Decode(&profile); err != nil {
		return profile, fmt.Errorf("error decoding document: %v", err)
	}

	return profile, nil
}

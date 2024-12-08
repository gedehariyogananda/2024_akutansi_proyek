package Repositories

import (
	"2024_akutansi_project/Models"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type (
	IDeviceTokenRepository interface {
		Create(ctx context.Context, payload Models.DeviceToken) (*Models.DeviceToken, error)
		Get(ctx context.Context, userID string) (*Models.DeviceToken, error)
	}

	DeviceTokenRepository struct {
		mongo *mongo.Client
	}
)

func DeviceTokenRepositoryProvider(mongo *mongo.Client) *DeviceTokenRepository {
	return &DeviceTokenRepository{mongo: mongo}
}

func (r *DeviceTokenRepository) Create(ctx context.Context, payload Models.DeviceToken) (*Models.DeviceToken, error) {
	collection := r.mongo.Database(os.Getenv("MONGO_DB_NAME")).Collection("device_token")

	// Use Upsert to ensure the document is either updated or inserted
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "user_id", Value: payload.UserID}}

	// Convert payload to a BSON document
	update := bson.D{{Key: "$set", Value: payload}}

	// Perform the upsert operation
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert document: %v", err)
	}

	fmt.Printf("Upserted document on user_id: %v\n", payload.UserID)

	return &payload, nil
}

func (r *DeviceTokenRepository) Get(ctx context.Context, userID string) (*Models.DeviceToken, error) {
	collection := r.mongo.Database(os.Getenv("MONGO_DB_NAME")).Collection("device_token")

	// Define the filter to find the document by user_id
	filter := bson.D{{Key: "user_id", Value: userID}}

	// Define a variable to hold the result
	var deviceToken Models.DeviceToken

	// Perform the find operation
	err := collection.FindOne(ctx, filter).Decode(&deviceToken)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no document found for user_id: %v", userID)
		}
		return nil, fmt.Errorf("failed to find document: %v", err)
	}

	fmt.Printf("Found document for user_id: %v\n", userID)

	return &deviceToken, nil
}

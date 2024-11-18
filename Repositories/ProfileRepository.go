package Repositories

import (
	"2024_akutansi_project/Models/Dto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type (
	IProfileRepository interface {
		IntegrateProfile(ctx context.Context, payload Dto.IntegratedProfile) (*Dto.IntegratedProfile, error)
		GetIntegratedProfile(ctx context.Context, id string) (*Dto.IntegratedProfile, error)
	}

	ProfileRepository struct {
		mongo *mongo.Client
	}
)

func ProfileRepositoryProvider(mongo *mongo.Client) *ProfileRepository {
	return &ProfileRepository{mongo: mongo}
}

// IntegrateProfile creates a new Profile document with a custom ID
func (r *ProfileRepository) IntegrateProfile(ctx context.Context, payload Dto.IntegratedProfile) (*Dto.IntegratedProfile, error) {
	collection := r.mongo.Database(os.Getenv("MONGO_DB_NAME")).Collection("profile")

	// Encrypt the token data
	payload.EncryptTokenData()

	// Use Upsert to ensure the document is either updated or inserted
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "user_id", Value: payload.UserID}}

	// Convert payload to a BSON document
	update := bson.D{{Key: "$set", Value: payload}}

	// Perform the upsert operation
	data, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert document: %v", err)
	}

	fmt.Printf("Upserted document with ID: %v\n", data.UpsertedID)

	return &payload, nil
}

func (r *ProfileRepository) GetIntegratedProfile(ctx context.Context, userID string) (*Dto.IntegratedProfile, error) {
	collection := r.mongo.Database(os.Getenv("MONGO_DB_NAME")).Collection("profile")

	// Create the filter to find the document by user_id
	filter := bson.D{{Key: "user_id", Value: userID}}

	var result struct {
		UserID      string  `bson:"user_id"`
		ShopeeToken *string `bson:"shopee_token,omitempty"`
		TiktokToken *string `bson:"tiktok_token,omitempty"`
	}

	// Find the document
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Map the result back to the DTO structure
	payload := &Dto.IntegratedProfile{
		UserID:      result.UserID,
		ShopeeToken: result.ShopeeToken,
		TiktokToken: result.TiktokToken,
	}

	payload.DecryptTokenData()

	return payload, nil
}

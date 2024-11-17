package Repositories

import (
	"2024_akutansi_project/Helper"
	"2024_akutansi_project/Models/Dto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type (
	IWebhookRepository interface {
		CreateWebhookProduct(ctx context.Context, payload Dto.CallbackProfile) (err error)
		GetWebhookByID(ctx context.Context, id string) (*Dto.CallbackProfile, error)
	}

	WebhookRepository struct {
		mongo *mongo.Client
	}
)

func WebhookRepositoryProvider(mongo *mongo.Client) *WebhookRepository {
	return &WebhookRepository{mongo: mongo}
}

// CreateWebhookProduct creates a new webhook document with a custom ID
func (r *WebhookRepository) CreateWebhookProduct(ctx context.Context, payload Dto.CallbackProfile) (err error) {
	collection := r.mongo.Database(os.Getenv("MONGO_DB_NAME")).Collection("webhook")

	// Marshal the payload into BSON (similar to JSON)
	document, err := bson.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Encrypt the entire BSON document
	encryptedDocument, err := Helper.EncryptData(document)
	if err != nil {
		return fmt.Errorf("failed to encrypt document: %v", err)
	}

	// Create a custom document structure with a custom _id field
	doc := bson.D{
		{Key: "_id", Value: payload.ID},         // Set the custom ID
		{Key: "data", Value: encryptedDocument}, // Store the encrypted document
	}

	// Use Upsert to ensure the document is either updated or inserted
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: payload.ID}}

	_, err = collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: doc}}, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert document: %v", err)
	}

	fmt.Printf("Upserted document with ID: %v\n", payload.ID)
	return nil
}

func (r *WebhookRepository) GetWebhookByID(ctx context.Context, id string) (*Dto.CallbackProfile, error) {
	collection := r.mongo.Database(os.Getenv("MONGO_DB_NAME")).Collection("webhook")

	var filter interface{}
	filter = bson.M{"_id": id}

	// Log the filter to ensure we're querying with the correct ID
	fmt.Printf("Query filter: %+v\n", filter)

	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("document with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to find document: %v", err)
	}

	// Log the result to see what is returned
	fmt.Printf("Found document: %+v\n", result)

	encryptedDocument, ok := result["data"].(string)
	if !ok {
		return nil, fmt.Errorf("data field is missing or not a string")
	}

	decryptedDocument, err := Helper.DecryptData(encryptedDocument)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt document: %v", err)
	}

	var payload Dto.CallbackProfile
	err = bson.Unmarshal(decryptedDocument, &payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal decrypted document: %v", err)
	}

	return &payload, nil
}

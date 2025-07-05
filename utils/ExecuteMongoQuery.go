package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoQuery struct {
	Operation  string                   `json:"operation"`
	Collection string                   `json:"collection"`
	Filter     map[string]interface{}   `json:"filter"`
	Database   string                   `json:"database,omitempty"`
	Projection map[string]interface{}   `json:"projection,omitempty"` // Optional, for find
	Update     map[string]interface{}   `json:"update,omitempty"`     // Optional, for update
	Documents  []map[string]interface{} `json:"documents,omitempty"`  // Optional, for insertMany
}

func ExecuteMongoQuery(client *mongo.Client, q MongoQuery) error {
	dbName := q.Database
	if dbName == "" {
		return fmt.Errorf("database name is required")
	}

	collection := client.Database(dbName).Collection(q.Collection)
	filter := q.Filter
	ctx := context.TODO()

	switch q.Operation {
	case "find":
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)
		var results []map[string]interface{}
		if err = cursor.All(ctx, &results); err != nil {
			return err
		}

		for _, result := range results {
			pretty, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(pretty))
		}

	case "count":
		n, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			return err
		}
		fmt.Printf("Count: %d\n", n)

	case "insert":
		res, err := collection.InsertOne(ctx, q.Documents[0])
		if err != nil {
			return err
		}
		fmt.Printf("Inserted ID: %v\n", res.InsertedID)

	case "insertMany":
		var docs []interface{}
		for _, d := range q.Documents {
			docs = append(docs, d)
		}
		res, err := collection.InsertMany(ctx, docs)
		if err != nil {
			return err
		}
		fmt.Printf("Inserted IDs: %v\n", res.InsertedIDs)

	case "update":
		res, err := collection.UpdateMany(ctx, filter, q.Update)
		if err != nil {
			return err
		}
		fmt.Printf("Updated count: %v\n", res.ModifiedCount)

	default:
		return fmt.Errorf("unsupported operation: %s", q.Operation)
	}
	return nil
}

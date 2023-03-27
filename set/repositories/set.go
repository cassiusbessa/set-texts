package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Set(db string, section string, content string) error {
	collection := Repo.Client.Database(db).Collection("texts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$set": bson.M{section: content}}
	_, err := collection.UpdateOne(ctx, bson.D{}, update)
	if err != nil {
		cancel()
		return err
	}
	return nil
}

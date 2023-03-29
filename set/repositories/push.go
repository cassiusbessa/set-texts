package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Push(db string, section string, content interface{}) error {
	collection := Repo.Client.Database(db).Collection("texts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$push": bson.M{section: content}}
	_, err := collection.UpdateOne(ctx, bson.D{}, update)
	if err != nil {
		cancel()
		return err
	}
	return nil
}

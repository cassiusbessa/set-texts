package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/cassiusbessa/db-texts/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Update(db string, id string, section string, entity entities.BaseEntity) (bool, error) {
	collection := Repo.Client.Database(db).Collection("texts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	mongoId, err := primitive.ObjectIDFromHex(id)
	entity.SetId(mongoId)
	if err != nil {
		cancel()
		return false, err
	}

	filter := bson.M{
		fmt.Sprintf("%s._id", section): mongoId,
	}
	arrayFilter := bson.M{"elem._id": mongoId}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("%s.$[elem]", section): entity,
		},
	}
	updateOpts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	result, err := collection.UpdateOne(ctx, filter, update, updateOpts)
	if err != nil {
		cancel()
		return false, err
	}
	if result.MatchedCount == 0 {
		cancel()
		return false, nil
	}

	defer cancel()
	return true, nil
}

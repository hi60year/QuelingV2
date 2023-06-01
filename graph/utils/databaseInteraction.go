package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/graph/constants"
)

func GetAllWithPagination[T any](dbClient *mongo.Client, ctx context.Context, collectionName string, filter any, pageNum int) ([]*T, error) {
	collection := dbClient.Database(constants.QuelingDatabaseName).Collection(collectionName)
	const pageSize = 20

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createdAt", -1}})
	findOptions.SetSkip(int64(pageNum * pageSize))
	findOptions.SetLimit(int64(pageSize))

	var results []*T

	if cur, err := collection.Find(ctx, filter, findOptions); err != nil {
		return nil, err
	} else {
		if err = cur.All(ctx, &results); err != nil {
			return nil, err
		} else {
			return results, err
		}
	}
}

func GetOneById[T any](dbClient *mongo.Client, ctx context.Context, collectionName string, id string) (*T, error) {
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return GetOne[T](dbClient, ctx, collectionName, bson.D{{"_id", bsonId}})
}

func GetOne[T any](dbClient *mongo.Client, ctx context.Context, collectionName string, filter any) (*T, error) {
	collection := dbClient.Database(constants.QuelingDatabaseName).Collection(collectionName)
	projection := createProjection(ctx, make(map[string]string))

	var result T

	if err :=
		collection.FindOne(ctx,
			filter,
			options.FindOne().SetProjection(projection)).Decode(&result); err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

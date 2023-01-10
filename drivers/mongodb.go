package drivers

import (
	"context"
	"errors"
	"github.com/hoitek-go/hoitek-cache/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct{}

var mongoConf *config.MongoDB
var mongoClient *mongo.Client

func initConf() {
	if mongoConf == nil {
		globalConfig := config.Global.(config.MongoDB)
		mongoConf = &globalConfig
	}
}

func Initialize() error {
	collections, err := mongoClient.Database(mongoConf.GetDatabase()).ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return err
	}
	found := false
	for _, collection := range collections {
		if collection == "caches" {
			found = true
			break
		}
	}
	if !found {
		err = mongoClient.Database(mongoConf.GetDatabase()).CreateCollection(context.Background(), "caches")
		if err != nil {
			return err
		}
	}
	return nil
}

func GetClient() *mongo.Client {
	if mongoClient == nil {
		conn, err := mongo.NewClient(options.Client().ApplyURI(mongoConf.GetURI()))
		if err != nil {
			log.Fatal(err)
		}
		err = conn.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		config.DB = &config.MongoDBClient{
			Client: conn,
		}
		mongoClient = conn
		err = Initialize()
		if err != nil {
			log.Fatal(err)
		}
	}
	return mongoClient
}

func (r MongoDB) Set(key string, value interface{}) error {
	initConf()
	mongoClient := GetClient()
	result := mongoClient.Database(mongoConf.Database).Collection("caches").FindOne(ctx, bson.D{{"key", key}})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			_, err := mongoClient.Database(mongoConf.Database).Collection("caches").UpdateOne(
				ctx,
				bson.D{{"key", key}},
				bson.D{{"$set", bson.D{{"value", value}}}},
				options.Update().SetUpsert(true),
			)
			if err != nil {
				return err
			}
			return nil
		}
		return result.Err()
	}
	_, err := mongoClient.Database(mongoConf.Database).Collection("caches").InsertOne(ctx, bson.D{{"key", key}})
	if err != nil {
		return err
	}
	return nil
}

func (r MongoDB) Get(key string) (interface{}, error) {
	initConf()
	mongoClient := GetClient()
	result := mongoClient.Database(mongoConf.Database).Collection("caches").FindOne(ctx, bson.D{{"key", key}})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("key not found")
		}
		return nil, result.Err()
	}
	var doc map[string]interface{}
	err := result.Decode(&doc)
	if err != nil {
		return nil, err
	}
	return doc["value"], nil
}

func (r MongoDB) Delete(key string) error {
	initConf()
	mongoClient := GetClient()
	_, err := mongoClient.Database(mongoConf.Database).Collection("caches").DeleteOne(ctx, bson.D{{"key", key}})
	if err != nil {
		return err
	}
	return nil
}

func (r MongoDB) Exists(key string) (bool, error) {
	initConf()
	doc, err := r.Get(key)
	if err != nil {
		return false, err
	}
	return doc != nil, nil
}

func (r MongoDB) Expire(key string, seconds int) error {
	initConf()
	return errors.New("not implemented")
}

func (r MongoDB) TTL(key string) (int, error) {
	initConf()
	return 0, errors.New("not implemented")
}

func (r MongoDB) Flush() error {
	initConf()
	mongoClient := GetClient()
	_, err := mongoClient.Database(mongoConf.Database).Collection("caches").DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}
	return nil
}

func (r MongoDB) Close() error {
	initConf()
	return errors.New("not implemented")
}

func (r MongoDB) Ping() error {
	initConf()
	return errors.New("not implemented")
}

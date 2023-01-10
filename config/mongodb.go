package config

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Host     string
	Port     int
	Username string
	Password string
	Timeout  int
	Database string
}

type MongoDBClient struct {
	Client *mongo.Client
}

type Cache struct {
	Key   string
	Value string
}

var DB *MongoDBClient

func (r MongoDB) GetHost() string {
	return r.Host
}

func (r MongoDB) GetPort() int {
	return r.Port
}

func (r MongoDB) GetTimeout() int {
	return r.Timeout
}

func (r MongoDB) GetUsername() string {
	return r.Username
}

func (r MongoDB) GetPassword() string {
	return r.Password
}

func (r MongoDB) GetAddress() string {
	return r.Host + ":" + fmt.Sprintf("%d", r.Port)
}

func (r MongoDB) GetURI() string {
	return "mongodb://" + r.GetUsername() + ":" + r.GetPassword() + "@" + r.GetAddress() + "/?directConnection=true&serverSelectionTimeoutMS=" + fmt.Sprintf("%d", r.GetTimeout())
}

func (r MongoDB) GetDatabase() string {
	return r.Database
}

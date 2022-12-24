package config_mongo

import (
	"context"
	configurations_yaml "go_sample/src/main/config/yaml"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MSG_MONGO_BEAN_INITIALIZING = "Initializing mongo database bean."
const MSG_MONGO_BEAN_PINGED = "Successfully connected and pinged."
const MSG_MONGO_BEAN_CLOSING_CONNECTION = "Closing connection of mongo database bean."
const MSG_ERROR_TO_CONNECT_TO_DATABASE = "Application has been failed to connect to mongo database. URI: %s"
const MSG_ERROR_TO_PING_DATABASE = "Application has been failed to ping mongo database. URI: %s"

const MONGO_URI_NAME = "MongoDB.URI"
const MONGO_DATABASE_NAME = "MongoDB.DatabaseName"

var MongoDatabase *mongo.Database = nil
var once sync.Once

func GetDatabaseBean() *mongo.Database {
	once.Do(func() { // <-- atomic, does not allow repeating

		if MongoDatabase == nil {
			MongoDatabase = getDatabaseConnection()
		} // <-- thread safe

	})
	return MongoDatabase
}

func CloseConnection() {
	log.Println(MSG_MONGO_BEAN_CLOSING_CONNECTION)
	MongoDatabase.Client().Disconnect(context.TODO())
}

func getDatabaseConnection() *mongo.Database {

	log.Println(MSG_MONGO_BEAN_INITIALIZING)

	databaseUri := configurations_yaml.GetBeanPropertyByName(MONGO_URI_NAME)
	mongoDatabaseName := configurations_yaml.GetBeanPropertyByName(MONGO_DATABASE_NAME)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatalf(MSG_ERROR_TO_CONNECT_TO_DATABASE, databaseUri)
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf(MSG_ERROR_TO_PING_DATABASE, databaseUri)
		panic(err)
	}
	log.Println(MSG_MONGO_BEAN_PINGED)

	return client.Database(mongoDatabaseName)

}

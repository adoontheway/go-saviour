package service

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Mongo *mongo.Client
}

var DB *Database

func SetConnect() *mongo.Client {
	uri := ""
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}
	return client
}

type mgo struct {
	database   string
	collection string
}

func NewMgo(database, collection string) *mgo {
	return &mgo{
		database,
		collection,
	}
}

func (m *mgo) FindOne(key string, value interface{}) *mongo.SingleResult {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	return singleResult
}

func (m *mgo) InsertOne(value interface{}) *mongo.InsertOneResult {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}

func (m *mgo) CollectionCount() (string, int64) {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	size, _ := collection.EstimatedDocumentCount(context.TODO())

	return collection.Name(), size
}

// query collection according to the given conditions
func (m *mgo) CollectionDocuments(skip, limit int64, sort int) *mongo.Cursor {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	SORT := bson.D{{"_id", sort}}
	filter := bson.D{{}}
	findOptions := options.Find().SetSort(SORT).SetLimit(limit).SetSkip(skip)
	temp, _ := collection.Find(context.Background(), filter, findOptions)

	return temp
}

// get collection create time and code
func (m *mgo) ParseId(result string) (time.Time, uint64) {
	temp1 := result[:8]
	timestamp, _ := strconv.ParseInt(temp1, 16, 64)
	datetime := time.Unix(timestamp, 0)
	temp2 := result[18:]
	count, _ := strconv.ParseUint(temp2, 16, 64)

	return datetime, count
}

// delete and query
func (m *mgo) DeleteAndFind(key string, value interface{}) (int64, *mongo.SingleResult) {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	deleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		log.Println("Error on delete")
	}

	return deleteResult.DeletedCount, singleResult
}

// delete
func (m *mgo) Delete(key string, value interface{}) int64 {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	filter := bson.D{{key, value}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		log.Println("Error on delete:", err)
	}

	return deleteResult.DeletedCount
}

// delete many
func (m *mgo) DeleteManagy(key string, value interface{}) int64 {
	client := DB.Mongo
	collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	filter := bson.D{{key, value}}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter, nil)
	if err != nil {
		log.Println("Error on delete many:", err)
	}

	return deleteResult.DeletedCount
}

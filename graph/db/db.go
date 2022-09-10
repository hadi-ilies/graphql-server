package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"
	"github.com/hadi-ilies/graphql-server/graph/model"
	"gopkg.in/mgo.v2/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE_NAME            = "graphql"
	DATABASE_COLLECTION_NAME = "videos"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

type Database struct {
	client *mongo.Client
}

func NewVideoRepository() VideoRepository {
	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MongoDB := os.Getenv("MongoDB")
	clientOptions := options.Client().ApplyURI(MongoDB).SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo Db Connected")
	return &Database{client: dbClient}
}

func (db *Database) Save(video *model.Video) {
	fmt.Println("SAVING")
	collection := db.client.Database(DATABASE_NAME).Collection(DATABASE_COLLECTION_NAME)
	collection.InsertOne(context.TODO(), video)
}

func (db *Database) FindAll() []*model.Video {
	collection := db.client.Database(DATABASE_NAME).Collection(DATABASE_COLLECTION_NAME)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var videos []*model.Video

	for cursor.Next(context.TODO()) {
		var video *model.Video
		err := cursor.Decode(&video)
		if err != nil {
			log.Fatal(err)
		}
		videos = append(videos, video)
	}
	return videos
}

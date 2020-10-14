package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fabiosebastiano/graphql-server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "graph-ql"
	COLLECTION = "videos"
)

func New() VideoRepository {

	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MONGODB := os.Getenv("MONGODB")

	clientOptions := options.Client().ApplyURI(MONGODB)
	clientOptions = clientOptions.SetMaxPoolSize(50)
	clientOptions = clientOptions.SetMaxConnIdleTime(10 * time.Second)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to MONGODB")

	return &database{
		client: dbClient,
	}
}
func (db *database) Save(video *model.Video) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Video salvato correttamente!")
}
func (db *database) FindAll() []*model.Video {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	defer cursor.Close(context.TODO())

	if err != nil {
		log.Fatal("Errore leggendo i video: ", err.Error())
	}

	var result []*model.Video
	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)

		if err != nil {
			log.Fatal("Errore leggendo il singolo video: ", err.Error())
		}
		result = append(result, v)
	}
	return result

}

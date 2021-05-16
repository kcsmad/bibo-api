package user

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type UserDAO struct {
	Client *mongo.Client
	Database string
}

func (dao *UserDAO) Disconnect() {
	err := dao.Client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}

func (dao *UserDAO) Connect() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_SERVER_URL")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	dao.Client = client
	dao.Database = os.Getenv("DB_NAME")
	collection := dao.Client.Database(dao.Database).Collection(os.Getenv("USER_COLLECTION_NAME"))
	return collection
}

func (dao *UserDAO) Insert(user *User) error {
	defer dao.Disconnect()

	collection := dao.Connect()
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

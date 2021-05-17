package book

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type BookDAO struct {
	Client *mongo.Client
	Database string
}

func (dao *BookDAO) Disconnect() {
	err := dao.Client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}

func (dao *BookDAO) Connect() *mongo.Collection {
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
	collection := dao.Client.Database(dao.Database).Collection(os.Getenv("BOOK_COLLECTION_NAME"))
	return collection
}

func (dao *BookDAO) InsertOne(book *Book) error {
	defer dao.Disconnect()

	collection := dao.Connect()
	_, err := collection.InsertOne(context.TODO(), book)
	return err
}

func (dao *BookDAO) InsertMany(books []*Book) error {
	defer dao.Disconnect()

	var interfaceBooks = make([]interface{}, 0, len(books))

	for _, book := range books {
		interfaceBooks = append(interfaceBooks, book)
	}

	collection := dao.Connect()
	_, err := collection.InsertMany(context.TODO(), interfaceBooks)

	return err
}

func (dao *BookDAO) GetAll() ([]*Book, error) {
	defer dao.Disconnect()

	var books []*Book

	collection := dao.Connect()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var book Book
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, &book)
	}

	return books, err
}

func (dao *BookDAO) GetByISBN(isbn int) (Book, error) {
	defer dao.Disconnect()

	var book Book

	collection := dao.Connect()
	err := collection.FindOne(context.TODO(), bson.M{"isbn": isbn}).Decode(&book)

	return book, err
}

func (dao *BookDAO) Update(book Book) error {
	defer dao.Disconnect()

	collection := dao.Connect()
	_, err := collection.UpdateByID(context.TODO(), book.Id, bson.D{{Key: "$set", Value: &book}})
	return err
}
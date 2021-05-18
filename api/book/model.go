package book

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id primitive.ObjectID `json:"-" bson:"_id"`
	Isbn int `json:"isbn" bson:"isbn"`
	Title string `json:"title" bson:"title"`
	AuthorName []string `json:"author_name" bson:"author_name"`
	ArtistName []string `json:"artist_name" bson:"artist_name"`
	Type string `json:"book_type" bson:"book_type"`
	PublisherName string `json:"publisher_name" bson:"publisher_name"`
	PublishYear int `json:"publish_year" bson:"publish_year"`
	Genres []string `json:"genres" bson:"genres"`
	Categories []string `json:"categories" bson:"categories"`
	Format string `json:"book_format" bson:"book_format"`
	CoverPrice string `json:"cover_price" bson:"cover_price"`
	TotalPages int `json:"total_pages" bson:"total_pages"`
	PaperType string `json:"paper_type" bson:"paper_type"`
	AdditionalInfo []string `json:"additional_info" bson:"additional_info"`
}
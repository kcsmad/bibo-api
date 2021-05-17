package book

import (
	. "bibo.api/api/rest"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

var dao = BookDAO{}

func CreateBook (c echo.Context) error {
	book := new(Book)
	if err := c.Bind(book); err != nil {
		return ResponseBadRequest(c, "Alguns dados estao faltando ou invalidos")
	}

	book.Id = primitive.NewObjectID()

	if err := dao.InsertOne(book); err != nil {
		return ResponseInternalError(c)
	}

	return ResponseCreated(c, book)
}

func FindAll(c echo.Context) error {
	books, err := dao.GetAll()

	if err != nil {
		return ResponseInternalError(c)
	}

	return ResponseSuccess(c, books)
}

func FindOne(c echo.Context) error {
	isbn, err := strconv.Atoi(c.Param("isbn"))

	if err != nil {
		return ResponseBadRequest(c, "ISBN invalido")
	}

	book, err := dao.GetByISBN(isbn)

	if err != nil {
		return ResponseInternalError(c)
	}

	return ResponseSuccess(c, book)
}

func Update(c echo.Context) error {
	newBook := new(Book)
	if err := c.Bind(newBook); err != nil {
		return ResponseBadRequest(c, "Alguns dados estao faltando ou invalidos")
	}

	book, err := dao.GetByISBN(newBook.Isbn)

	if err != nil {
		return ResponseInternalError(c)
	}

	//if (book == Book{}) {  '== not an operator from Book'
	//	return ResponseNotFound(c)
	//}

	newBook.Id = book.Id

	err = dao.Update(*newBook)

	if err != nil {
		return ResponseInternalError(c)
	}

	return ResponseSuccess(c, newBook)
}
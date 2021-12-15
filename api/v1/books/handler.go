package books

import (
	. "bibo.api/api/v1/rest"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
)

type Handler struct {
	DAO DAO
}

var booksHandlerLogger = log.WithFields(log.Fields{
	"[domain]": "[Books Handler]",
	"[env]":    "[" + os.Getenv("APP_ENV") + "]",
})

func (h *Handler) FindAll(c echo.Context) error {
	books, err := h.DAO.GetAll()

	if err != nil {
		booksHandlerLogger.WithFields(log.Fields{
			"[method]": "[FindAll]",
			"[action]": "[Get All From Database]",
		}).Fatal(err.Error())

		return ResponseInternalError(c)
	}

	for _, book := range books {
		if err := h.buildBook(book, c); err != nil {
			return err
		}
	}

	return ResponseSuccess(c, books)
}

func (h *Handler) buildBook(book *Book, c echo.Context) error {

	covers, err := h.DAO.GetBookCovers(book.ISBN)

	if err != nil {
		booksHandlerLogger.WithFields(log.Fields{
			"[method]": "[Build Book]",
			"[action]": "[Get Covers from Database]",
		}).Fatal(err.Error())

		return ResponseInternalError(c)
	}

	bookType, err := h.DAO.GetBookType(book.BookTypeId)

	if err != nil {
		booksHandlerLogger.WithFields(log.Fields{
			"[method]": "[Build Book]",
			"[action]": "[Get Book Type from Database]",
		}).Fatal(err.Error())

		return ResponseInternalError(c)
	}

	paperType, err := h.DAO.GetPaperType(book.PaperTypeId)

	if err != nil {
		booksHandlerLogger.WithFields(log.Fields{
			"[method]": "[Build Book]",
			"[action]": "[Get Paper Type from Database]",
		}).Fatal(err.Error())

		return ResponseInternalError(c)
	}

	authors, err := h.DAO.GetAuthors(book.ISBN)

	if err != nil {
		booksHandlerLogger.WithFields(log.Fields{
			"[method]": "[Build Book]",
			"[action]": "[Get Authors from Database]",
		}).Fatal(err.Error())

		return ResponseInternalError(c)
	}

	book.Covers = covers
	book.BookType = bookType
	book.PaperType = paperType
	book.Authors = authors

	return nil
}

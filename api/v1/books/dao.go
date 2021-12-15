package books

import (
	. "bibo.api/api/v1/authors"
	"bibo.api/api/v1/db"
	"os"
)

type DAO struct {
	Repo *db.Repository
}

func (dao *DAO) GetAll() ([]*Book, error) {
	query := `SELECT * FROM ` + os.Getenv("BOOK_TABLE_NAME") + ` WHERE is_active = 1;`

	rows, err := dao.Repo.Client.Query(query)

	if err != nil {
		return nil, err
	}

	var books []*Book

	for rows.Next() {
		var book Book

		if err := rows.Scan(
			&book.ISBN, &book.ISBN10, &book.Title, &book.Slug,
			&book.BookTypeId, &book.BookFormat, &book.CoverPrice, &book.PaperTypeId,
			&book.IsPublished, &book.IsActive, &book.TotalPages,
		); err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func (dao *DAO) GetBookCovers(isbn int) ([]string, error) {
	query := `SELECT source FROM ` + os.Getenv("BOOK_COVERS_TABLE_NAME") + " WHERE isbn = ?;"

	rows, err := dao.Repo.Client.Query(query, isbn)

	if err != nil {
		return nil, err
	}

	var covers []string

	for rows.Next() {
		var cover string

		if err := rows.Scan(&cover); err != nil {
			return nil, err
		}

		covers = append(covers, cover)
	}

	return covers, nil
}

func (dao *DAO) GetBookType(id int) (string, error) {
	query := `SELECT description FROM ` + os.Getenv("BOOK_TYPES_TABLE_NAME") + " WHERE id = ?;"

	rows, err := dao.Repo.Client.Query(query, id)

	if err != nil {
		return "", err
	}

	var bookType string

	for rows.Next() {
		if err := rows.Scan(&bookType); err != nil {
			return "", err
		}
	}

	return bookType, nil
}

func (dao *DAO) GetPaperType(id int) (string, error) {
	query := `SELECT description FROM ` + os.Getenv("BOOK_PAPER_TYPE_TABLE_NAME") + " WHERE id = ?;"

	rows, err := dao.Repo.Client.Query(query, id)

	if err != nil {
		return "", err
	}

	var paperType string

	for rows.Next() {
		if err := rows.Scan(&paperType); err != nil {
			return "", err
		}
	}

	return paperType, nil
}

func (dao *DAO) GetAuthors(isbn int) ([]*Author, error) {
	query := `SELECT a.name, a.slug FROM ` + os.Getenv("BOOK_AUTHORS_TABLE_NAME") +
		" ba INNER JOIN " + os.Getenv("AUTHOR_TABLE_NAME") + " a ON ba.author_id = a.id" +
		" WHERE ba.isbn = ?"

	rows, err := dao.Repo.Client.Query(query, isbn)

	if err != nil {
		return nil, err
	}

	var authors []*Author

	for rows.Next() {
		var author Author
		if err := rows.Scan(&author.Name, &author.Slug); err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return authors, nil
}

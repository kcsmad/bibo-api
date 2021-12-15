package books

import . "bibo.api/api/v1/authors"

type Book struct {
	ISBN        int       `json:"isbn" db:"isbn"`
	ISBN10      int       `json:"isbn_10" db:"isbn_10"`
	Title       string    `json:"title" db:"title"`
	Slug        string    `json:"slug" db:"slug"`
	BookTypeId  int       `json:"-" db:"book_type_id"`
	BookType    string    `json:"book_type" db:"-"`
	PaperTypeId int       `json:"-" db:"paper_type_id"`
	PaperType   string    `json:"paper_type" db:"-"`
	BookFormat  string    `json:"book_format" db:"book_format"`
	CoverPrice  string    `json:"cover_price" db:"cover_price"`
	IsPublished bool      `json:"is_published" db:"is_published"`
	IsActive    bool      `json:"-" db:"is_active"`
	TotalPages  int       `json:"total_pages" db:"total_pages"`
	Covers      []string  `json:"covers" db:"-"`
	Authors     []*Author `json:"authors" db:"-"`
}

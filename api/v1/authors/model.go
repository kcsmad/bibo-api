package authors

type Author struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" db:"name"`
	Slug string `json:"slug" db:"slug"`
}

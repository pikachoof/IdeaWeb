package models

type Author struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthorRepositoryInterface interface {
	GetAuthorByID(authorID uint) (Author, error)
	GetAllAuthors() ([]Author, error)
	AddAuthor(firstName string, lastName string) error
	RemoveAuthorByID(authorID uint) error
	RemoveAuthorByCredentials(firstName string, lastName string) error
}

package model

type Book struct {
	ID       int64
	AuthorID int
	Title    string
	ISBN     string
	Subject  string
}

type Author struct {
	ID        int
	FirstName string
	LastName  string
}

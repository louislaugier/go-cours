package article

type Article struct {
	Title    string
	Content  string
	Category string
	Tags     []string
	Author   *User
}

type User struct {
	Username  string
	FirstName string
	LastName  string
}

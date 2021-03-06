package blog

import (
	"strconv"

	"github.com/louislaugier/go-cours/example/article"
)

// Blog export
type Blog interface {
	String() string
	GetPostByID(id string) *article.Article
	GetLatestPosts(limit int) []*article.Article
}

// Error export
type Error struct {
	StatusCode  int
	Message     string
	Validations []*Validation
}

// Validation export
type Validation struct {
	Message string
	Path    []string
}

// Error export
func (e *Error) Error() string {
	txt := e.Message + "(" + strconv.Itoa(e.StatusCode) + ")"
	if e.StatusCode != 0 {
		txt += "(" + strconv.Itoa(e.StatusCode) + ")"
	}
	return txt
}

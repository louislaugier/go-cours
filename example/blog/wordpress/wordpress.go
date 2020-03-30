package wordpress

import (
	"fmt"
	"time"

	"github.com/louisl98/go-rappels/example/article"
	"github.com/louisl98/go-rappels/example/blog"
)

// WordPress export
type WordPress struct {
	URL      string
	Username string
	Password string
}

// String export
func (wp *WordPress) String() string {
	return "WordPress"
}

// GetPostByID export
func (wp *WordPress) GetPostByID(id string) *article.Article {
	time.Sleep(2 * time.Second)
	fmt.Println("Getting posts by ID from a WordPress blog...")
	return &article.Article{
		Title: "This comes from a WordPress blog",
	}
}

// GetLatestPosts export
func (wp *WordPress) GetLatestPosts(limit int) []*article.Article {
	time.Sleep(2 * time.Second)
	fmt.Println("Getting posts from a WordPress blog...")
	return []*article.Article{
		{
			Title: "Trop bien WordPress",
		},
	}
}

// New export
func New(wp *WordPress) (*WordPress, error) {
	if wp == nil {
		return nil, &blog.Error{
			Message: "Options must not be nil",
		}
	}
	if wp.URL == "" {
		return nil, fmt.Errorf("No URL for WordPress")
	}
	if wp.Username == "" {
		return nil, &blog.Error{
			StatusCode: 500,
			Message:    "Username not set",
		}
	}
	return wp, nil
}

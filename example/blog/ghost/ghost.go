package ghost

import (
	"fmt"
	"time"

	"github.com/louisl98/go-rappels/example/article"
	blog "github.com/louisl98/go-rappels/example/blog"
)

// Ghost export
type Ghost struct {
	URL      string
	Username string
	Password string
}

// export String
func (gh *Ghost) String() string {
	return "Ghost"
}

// GetPostByID export
func (gh *Ghost) GetPostByID(id string) *article.Article {
	time.Sleep(2 * time.Second)
	fmt.Println("Getting posts by ID from a Ghost blog...")
	return &article.Article{
		Title: "This comes from a WordPress blog",
	}
}

// GetLatestPosts export
func (gh *Ghost) GetLatestPosts(limit int) []*article.Article {
	time.Sleep(3 * time.Second)
	fmt.Println("Getting posts from a Ghost blog...")
	return []*article.Article{
		{
			Title: "Trop bien Ghost",
		},
	}
}

// New export
func New(wp *Ghost) (*Ghost, error) {
	if wp == nil {
		wp = &Ghost{}
	}
	if wp.URL == "" {
		return nil, fmt.Errorf("No URL for Ghost")
	}
	if wp.Username == "" {
		return nil, &blog.Error{
			StatusCode: 500,
			Message:    "Username not set",
		}
	}
	return wp, nil
}

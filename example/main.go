package main

import (
	"fmt"
	"sync"

	"github.com/louisl98/go-rappels/example/article"
	"github.com/louisl98/go-rappels/example/blog"
	"github.com/louisl98/go-rappels/example/blog/ghost"
	"github.com/louisl98/go-rappels/example/blog/wordpress"
)

type feed struct {
	Articles []*article.Article
	Blogs    []blog.Blog
}

type details struct {
	Name string
	URL  string
}

func main() {
	myFeed := &feed{
		Articles: []*article.Article{},
		Blogs: []blog.Blog{
			&ghost.Ghost{},
			&wordpress.WordPress{},
		},
	}
	wg := sync.WaitGroup{}
	fmt.Println(myFeed)
	for _, b := range myFeed.Blogs {
		wg.Add(1)
		// myFeed.Articles = append(myFeed.Articles, b.GetLatestPosts(5)...) // ... = appliquer chaque entrée
		go func(wg *sync.WaitGroup, b blog.Blog) {
			defer wg.Done()
			b.GetLatestPosts(5)
		}(&wg, b)
	}
	wg.Wait()
	fmt.Println("Done !")
}

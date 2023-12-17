package main

import "fmt"

type Article struct {
	Title, Description string
	IsActive           bool
}

func main() {
	var article *Article = &Article{}
	// var article *Article = new(Article)
	var article2 *Article = article
	// article2 := article

	// article.Description = "Berita 1"

	fmt.Println(article)
	fmt.Println(article2)

	article2.Description = "Isi Berita"
	article2.IsActive = true
	fmt.Println(article)
	fmt.Println(article2)
}

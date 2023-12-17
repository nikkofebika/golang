package main

import "fmt"

type Article struct {
	Title, Description string
	IsActive           bool
}

func main() {
	// var article Article = Article{"Berita 1", "Isi Berita", true}
	// var article2 *Article = &article

	// article2.Description = "Edit Isi Berita"
	// fmt.Println(article)
	// fmt.Println(article2)

	// article2 = &Article{"Berita 2", "Kebakarannnnn", false}

	// fmt.Println(article)
	// fmt.Println(article2)

	// asterisk, ubah semua data yang berkaitan
	var article Article = Article{"Berita 1", "Isi Berita", true}
	var article2 *Article = &article
	var article3 *Article = article2

	article2.Description = "Edit Isi Berita"
	fmt.Println(article)
	fmt.Println(article2)
	fmt.Println(article3)

	// article2 = &Article{"Berita 2", "Kebakarannnnn", false}
	*article2 = Article{"Berita 2", "Kebakarannnnn", false}

	fmt.Println(article)
	fmt.Println(article2)
	fmt.Println(article3)
}

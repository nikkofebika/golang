package main

import "fmt"

type Article struct {
	Title, Description string
	IsActive           bool
}

func main() {
	var article Article = Article{Title: "Berita 1"}
	// article2 := article // pass by value, by default golang akan men-copy data dari variabel yang dituju
	// article2 := &article // pointer/pass by reference, data dari var article2 akan mengacu ke var article. apabila article2 diubah datanya, maka var article akan ikut berubah datanya
	var article2 *Article = &article

	fmt.Println(article)
	fmt.Println(article2)

	article2.Description = "Isi Berita"
	article2.IsActive = true
	fmt.Println(article)
	fmt.Println(article2)
}

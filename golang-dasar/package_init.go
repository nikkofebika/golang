package main

// gunakan _ untuk memanggil init tanpa memanggil function2 nya
import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/services"
)

func main() {
	fmt.Println(database.GetConnection())
}

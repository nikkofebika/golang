package database

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	fmt.Println("Connection success ", db.Stats())
}

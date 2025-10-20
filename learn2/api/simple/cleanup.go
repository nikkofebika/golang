package simple

import "fmt"

type File struct {
	Name string
}

func (file *File) Close() {
	fmt.Println("Close file", file.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}

type Connection struct {
	*File
}

func (connection *Connection) Close() {
	fmt.Println("Close connection File Name", connection.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}

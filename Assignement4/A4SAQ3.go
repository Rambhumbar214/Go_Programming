package main

import "fmt"

type Author struct {
	Name  string
	Book  string
	Email string
}

func (a Author) Show() {
	fmt.Println("Author Name:", a.Name)
	fmt.Println("Book Title:", a.Book)
	fmt.Println("Email:", a.Email)
}

func main() {
	author := Author{Name: "John Doe", Book: "Go Programming Basics", Email: "john@example.com"}
	author.Show()
}

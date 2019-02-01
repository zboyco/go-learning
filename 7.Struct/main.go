package main

import "fmt"

type Book struct {
	id      int
	title   string
	context string
}

func main() {
	book1 := Book{id: 1, title: "book1", context: "context for book1"}
	book1.title = "new book1"

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 id : %d\n", book1.id)
	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 context : %s\n", book1.context)

	UpdateBookId(book1)

	fmt.Printf("Book 1 updated id : %d\n", book1.id)

	var struct_pointer *Book
	struct_pointer = &book1

	/* 使用指针打印 Book1 信息 */
	fmt.Printf("Book 1 id : %d\n", struct_pointer.id)
	fmt.Printf("Book 1 title : %s\n", struct_pointer.title)
	fmt.Printf("Book 1 context : %s\n", struct_pointer.context)

	UpdateBookIdByPointer(&book1)

	fmt.Printf("Book 1 updated id : %d\n", book1.id)
}

func UpdateBookId(book Book) {
	book.id++
}

func UpdateBookIdByPointer(book *Book) {
	book.id++
}

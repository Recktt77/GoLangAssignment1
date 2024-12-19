package main

import {
	"fmt"

}

func main() {
	function := ""
	id := ""
	book := ""
	title := ""
	library := Library{}
	library.init()
	for true {
		println("chouse the function(addbook, borrow, return and listbooks or exit)")
		fmt.Scan(&function)
		if function == "addbook" {
			fmt.Println("write id(srting), book(string), title(string)")
			fmt.Scanf("%s %s %s", &id, &book, &title)
			if id == "" || book == "" || title == "" {
				fmt.Println("Do not skip empty")
				return
			}
			library.AddBook(id, book, title)
		} else if function == "borrow" {
			fmt.Println("write id(srting)")
			fmt.Scanf(id)
			if id == "" {
				fmt.Println("Do not skip empty")
				return
			}
			library.BorrowBook(id)
		} else if function == "return" {
			fmt.Println("write id(srting)")
			fmt.Scanf(id)
			if id == "" {
				fmt.Println("Do not skip empty")
				return
			}
			library.ReturnBook(id)
		} else if function == "listbooks" {
			library.ListBooks()
		} else if function == "exit" {
			break
		}
	}
}

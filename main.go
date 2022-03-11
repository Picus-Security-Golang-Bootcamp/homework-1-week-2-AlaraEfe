package main

import (
	"fmt"
	"os"
	"strings"
)

type books struct {
	bookID   int64
	bookName string
}

func showBooks() []books {

	book1 := books{
		bookID:   1,
		bookName: "Guns, Germs, and Steel",
	}

	book2 := books{
		bookID:   2,
		bookName: "Escape From Freedom",
	}

	book3 := books{
		bookID:   3,
		bookName: "The Grapes of Wrath",
	}

	book4 := books{
		bookID:   4,
		bookName: "Blindness",
	}

	book5 := books{
		bookID:   5,
		bookName: "Crime and Punishment",
	}

	book6 := books{
		bookID:   6,
		bookName: "Letter to a Child Never Born",
	}

	return []books{book1, book2, book3, book4, book5, book6}

}

func main() {

	args := os.Args[1:]

	booksList := showBooks()

	if args[0] == "list" {

		if len(args) == 1 {

			fmt.Println("The list of the Book Archive shown below;")

			for i := 0; i < len(booksList); i++ {
				fmt.Println(booksList[i])
			}

		} else {
			fmt.Println("Only the 'list' command should be entered for seeing the list of the Book Archive")
		}

	} else if args[0] == "search" {

		if len(args) > 1 {

			var userBookNameInput string
			userBookNameInput = strings.Join(args[1:], " ")

			var bookExistInList bool
			bookExistInList = false

			for i := 0; i < len(booksList); i++ {

				if strings.ToLower(userBookNameInput) == strings.ToLower(booksList[i].bookName) {
					s := fmt.Sprintf("The book that searched for is '%s'", booksList[i].bookName)
					print(s)
					//fmt.Println(booksList[i].bookName)
					bookExistInList = true
				}

			}

			if bookExistInList == false {
				fmt.Println("The book does not exist in the list")
			}

		} else {
			fmt.Println("Please specify the book name after the 'search' command as 'search <the name of the book>'")
		}

	} else {
		fmt.Println("The 'list' or 'search' commands should be used for seeing the list of the Book Archive or searching the specific book in the Book Archive.")
	}

}

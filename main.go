// @Author: Adrian Rincon Jimenez
// GO CLI
// This is my main.go file. This is where you will be running the program from. Run the file with go run main.go
// and then press on the keyboard which options you would like to pick.

// I used documentation to convert what I had previously from python to golang
// https://gosamples.dev/read-user-input/
// I also looked at this youtube video to get somewhat of an overview
// https://www.youtube.com/watch?v=yyUHQIec83I

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// make book struct with ID, bookName, AuthorName, Status
type Book struct {
	ID         int
	BookName   string
	AuthorName string
	Status     string
}

// a list of books that are in the library
var bookList = []Book{
	{1, "The Brief and Wondrous Life of Oscar Wao", "Junot Diaz", "Available"},
	{2, "Roots", "Alex Haley", "Available"},
	{3, "I Have No Mouth, and I Must Scream", "Harlan Ellison", "Available"},
	{4, "To Kill a Mockingbird", "Harper Lee", "Available"},
	{5, "The Great Gatsby", "F. Scott Fitzgerald", "Available"},
}

// This function allows you to return books. Type in their ID and then they are found.
// Checks to see if books are already returned. If they are then user cannot return book.
func returnBook() {
	fmt.Print("Enter the ID of the book you would like to return (type 0 to return to the menu): ")
	input := userInput()
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	if id == 0 {
		return
	}
	for i, book := range bookList {
		if book.ID == id && book.Status == "Checked-out" {
			bookList[i].Status = "Available"
			fmt.Println("You have returned the book successfully.")
			return
		} else if book.ID == id && book.Status == "Available" {
			fmt.Println("That book has already been returned.")
			return
		}
	}
	fmt.Println("No book found with that ID or it is not checked out.")
}

// This function lets you check out books. when ID is typed, then it checks to see if they book is available or not
func checkOutBook() {
	fmt.Print("Enter the ID of the book you would like to check out (type 0 to return to the menu): ")
	input := userInput()
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	if id == 0 {
		return
	}
	for i, book := range bookList {
		if book.ID == id && book.Status == "Available" {
			bookList[i].Status = "Checked-out"
			fmt.Println("You have checked out the book successfully.")
			return
		} else if book.ID == id && book.Status == "Checked-out" {
			fmt.Println("That book is already checked out.")
			return
		}
	}
	fmt.Println("No available book found with that ID.")
}

// this function prints out the menu options.
func printMenu() {
	fmt.Println("Please select from the following options...")
	fmt.Println("1. Show book list")
	fmt.Println("2. Search for book by ID number")
	fmt.Println("3. Return book")
	fmt.Println("4. Check out book")
	fmt.Println("5. Exit\n")
}

// this function shows the list of books with their book ID, Title, Author, and Status
func showBooks() {
	fmt.Println("Book list:")
	for _, book := range bookList {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.BookName, book.AuthorName, book.Status)
	}
}

// This allows you to search for a book and display all of its relevant information when you input the book's ID
func idSearch() {
	for {
		fmt.Print("Enter the ID of the book you would like to search for (type 0 to return to the menu): ")
		input := userInput()
		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if id == 0 {
			return
		}
		found := false
		for _, book := range bookList {
			if book.ID == id {
				fmt.Printf("Book found:\nID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.BookName, book.AuthorName, book.Status)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("No book found with that ID.")
		}
	}
}

// extra function that allows me to read user input.
// got this code from here:
// https://zetcode.com/golang/readinput/
func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

// main function runs all the code in a loop
func main() {
	fmt.Println("Welcome to Rincon Library!\n")
	for {
		printMenu()
		fmt.Print("Enter your choice: ")
		input := userInput()
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}
		switch choice {
		case 1:
			showBooks()
		case 2:
			idSearch()
		case 3:
			returnBook()
		case 4:
			checkOutBook()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a number from the list.")
		}
	}
}

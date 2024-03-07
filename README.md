# Rincon Library CLI

## Introduction
This command-line interface (CLI) program, written in Go, serves as a simple library management system. It allows users to perform various operations such as viewing the list of available books, searching for books by ID, checking out books, and returning books.

## Usage
To run the program, execute the 'main.go' file using the following command:

go run main.go

Once running, choose from the given prompts to navigate through the menu and perform desired actions.

## Features
### Book Struct
The program defines a 'Book' struct with the following attributes:
- 'ID': Unique identifier for each book.
- 'BookName': Title of the book.
- 'AuthorName': Name of the author.
- 'Status': Current status of the book (Available or Checked-out).

### Menu Options
Main menu has the following options:
1. **Show book list**: Display the list of available books with their details.
2. **Search for book by ID number**: Search for a specific book by its ID and view its details.
3. **Return book**: Return a checked-out book by entering its ID.
4. **Check out book**: Check out an available book by entering its ID.
5. **Exit**: Exit Program.

### Functionality
- Users can view the list of books, search for a book by its ID, check out available books, and return checked-out books.
- Input validation is implemented to handle invalid user input.
- The program runs in a loop, allowing users to perform multiple operations without exiting.

## Author
Adrian Rincon Jimenez.

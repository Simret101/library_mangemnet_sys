package controllers

import (
	"ans/models"
	"ans/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Controller struct {
	libraryService *services.LibService
}

func NewController() *Controller {
	return &Controller{
		libraryService: services.NewService(),
	}
}

// in struct it is always necessary to make instance of the struct inorder for the method to manipulate it ,preferalble that it is pointer instead of value
func (ans *Controller) Run() {
	for {
		ans.menu()
		option := ans.takeInput()
		ans.inputOptions(option)
	}
}

// To display the menu
func (ans *Controller) menu() {
	fmt.Println("-----------Welcome to Library Management System!---------")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Exit")
}

// To take input
func (c *Controller) takeInput() int {
	var options int

	fmt.Print("Please choose: ")
	option, _ := readLine()
	options, _ = strconv.Atoi(option)
	return options
}

// To choose an input
func (ans *Controller) inputOptions(options int) {
	switch options {
	case 1:
		ans.addBook()
	case 2:
		ans.removeBook()
	case 3:
		ans.borrowBook()
	case 4:
		ans.returnBook()
	case 5:
		ans.listAvailableBooks()
	case 6:
		ans.listBorrowedBooks()
	case 7:
		fmt.Println("Exit")
		return
	default:
		fmt.Println("Please try again.")
	}
}

// To add a book
func (ans *Controller) addBook() {
	//this specific book is part of the go model package it indicates that
	var book models.Book

	for {
		fmt.Print("Enter book ID: ")
		bookID, _ := readLine()
		if isValidID(bookID) {
			book.ID, _ = strconv.Atoi(bookID)
			break
		} else {
			fmt.Println("Book ID must contain only numbers.")
		}
	}

	for {
		fmt.Print("Enter book title: ")
		bookTitle, _ := readLine()
		if isValidName(bookTitle) {
			book.Title = bookTitle
			break
		} else {
			fmt.Println("Book title must contain only letters.")
		}
	}

	for {
		fmt.Print("Enter book author: ")
		bookAuthor, _ := readLine()
		if isValidName(bookAuthor) {
			book.Author = bookAuthor
			break
		} else {
			fmt.Println("Book author must contain only letters.")
		}
	}

	book.Status = "Available"
	ans.libraryService.AddBook(book)
	fmt.Println("Book added successfully.")
}

// To remove a book
func (ans *Controller) removeBook() {
	var bookID int

	for {
		fmt.Print("Enter book ID to remove: ")
		bookIDStr, _ := readLine()
		if isValidID(bookIDStr) {
			bookID, _ = strconv.Atoi(bookIDStr)
			break
		} else {
			fmt.Println("Book ID must contain only numbers.")
		}
	}

	ans.libraryService.RemoveBook(bookID)
	fmt.Println("Book removed successfully.")
}

// To borrow a book
func (ans *Controller) borrowBook() {
	var bookID, memberID int

	for {
		fmt.Print("Enter book ID to borrow: ")
		bookIDStr, _ := readLine()
		if isValidID(bookIDStr) {
			bookID, _ = strconv.Atoi(bookIDStr)
			break
		} else {
			fmt.Println("Book ID must contain only numbers.")
		}
	}

	for {
		fmt.Print("Enter member ID : ")
		memberIDStr, _ := readLine()
		if isValidID(memberIDStr) {
			memberID, _ = strconv.Atoi(memberIDStr)
			break
		} else {
			fmt.Println("Member ID must contain only numbers.")
		}
	}

	err := ans.libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

// To return a book borrowed
func (ans *Controller) returnBook() {
	var bookID, memberID int

	for {
		fmt.Print("Enter book ID to return: ")
		bookIDStr, _ := readLine()
		if isValidID(bookIDStr) {
			bookID, _ = strconv.Atoi(bookIDStr)
			break
		} else {
			fmt.Println("Book ID must contain only numbers.")
		}
	}

	for {
		fmt.Print("Enter member ID: ")
		memberIDStr, _ := readLine()
		if isValidID(memberIDStr) {
			memberID, _ = strconv.Atoi(memberIDStr)
			break
		} else {
			fmt.Println("Member ID must contain only numbers.")
		}
	}

	err := ans.libraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

// To display list of available books
func (ans *Controller) listAvailableBooks() {
	books := ans.libraryService.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		//%d stands for "decimal integer"
		//%s stands for "string"
		//%f - Floating-point number
		//%c - Single character
		fmt.Printf("Book ID: %d, Book Title: %s, Book Author: %s\n", book.ID, book.Title, book.Author)
	}
}

// To display list of borrowed books
func (ans *Controller) listBorrowedBooks() {
	var memberID int

	for {
		fmt.Print("Enter member ID (numbers only): ")
		memberIDStr, _ := readLine()
		if isValidID(memberIDStr) {
			memberID, _ = strconv.Atoi(memberIDStr)
			break
		} else {
			fmt.Println("Member ID must contain only numbers.")
		}
	}

	books := ans.libraryService.ListBorrowedBooks(memberID)
	fmt.Printf("Books borrowed by member %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

// To check input eror
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(ans), nil
}

// To validate the input name
func isValidName(name string) bool {
	if name == "" {
		return false
	}

	for _, s := range name {
		if !unicode.IsLetter(s) {
			return false
		}
	}

	return true
}

func isValidID(id string) bool {
	if id == "" {
		return false
	}
	//Package unicode provides data and functions to test some properties of Unicode code points
	for _, s := range id {
		if !unicode.IsDigit(s) {
			return false
		}
	}

	return true
}

/*writing effective comments to improve code structure
- Begin each comment with a clear explanation of the purpose or intent of the code block.
- If the code contains complex logic, algorithms, or business rules, use comments to explain the rationale and steps involved.
- Call out any important assumptions the code makes, as well as how it handles edge cases or unexpected inputs.
- For functions, include comments that explain what each parameter represents and what the function returns.
- Avoid lengthy, wordy comments
- Use Markdown syntax*/

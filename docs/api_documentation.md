# Library Management System Documentation

## Overview
The Library Management System is a console-based application built in Go. It provides functionality to manage books and members in a library.

## Structs
1. **Book**:
   - ID(int): Unique identifier for the book.
   - Title(string): Title of the book.
   - Author(string): Author of the book.
   - Status(string): Current status of the book ("Available" or "Borrowed").

2. **Member**:
   - ID(int): Unique identifier for the member.
   - Name(string): Name of the member.
   - BorrowedBooks ([]Book): Slice of books borrowed by the member.

## Interfaces
1. **LibraryManager**:
   - AddBook(book Book): Adds a new book to the library.
   - RemoveBook(bookID int): Removes a book from the library by its ID.
   - BorrowBook(bookID int, memberID int) error: Allows a member to borrow a book if it is available.
   - ReturnBook(bookID int, memberID int) error: Allows a member to return a borrowed book.
   - ListAvailableBooks() []Book: Lists all available books in the library.
   - ListBorrowedBooks(memberID int) []Book: Lists all books borrowed by a specific member.

## Implementation
1. **LibraryService**:
   - Implements the "LibraryManager" interface.
   - Stores the books and members in separate maps.
   - Provides methods to perform various library management operations.

2. **LibraryController**:
   - Handles the console input and invokes the appropriate service methods.
   - Provides a menu interface for the user to interact with the library management system.

## Folder Structure
- main.go: Entry point of the application.
- controllers/library_controller.go: Handles console input and invokes the appropriate service methods.
- models/book.go: Defines the "Book" struct.
- models/member.go: Defines the "Member"struct.
- services/library_service.go: Contains business logic and data manipulation functions.
- docs/documentation.md: Contains system documentation and other related information.
- go.mod: Defines the module and its dependencies.

## Usage
1. Run the application using the "go run main.go" command.



package services

import (
	"ans/models"
	"errors"
)

type LibService struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewService() *LibService {
	return &LibService{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (ans *LibService) AddBook(book models.Book) {
	ans.books[book.ID] = book
}

func (ans *LibService) RemoveBook(bookID int) {
	delete(ans.books, bookID)
}

func (ans *LibService) BorrowBook(bookID, memberID int) error {
	book, ok := ans.books[bookID]
	if !ok {
		return errors.New(" The Book is not found")
	}

	if book.Status != "Available" {
		return errors.New("The Book is not available")
	}

	member, ok := ans.members[memberID]
	if !ok {
		return errors.New("The member is not found")
	}

	book.Status = "Borrowed"
	ans.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	ans.members[memberID] = member
	return nil
}

func (ans *LibService) ReturnBook(bookID, memberID int) error {
	book, ok := ans.books[bookID]
	if !ok {
		return errors.New("The book is not found")
	}

	if book.Status != "Borrowed" {
		return errors.New("The book is not borrowed")
	}

	member, ok := ans.members[memberID]
	if !ok {
		return errors.New("The member is not found")
	}

	book.Status = "Available"
	ans.books[bookID] = book

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	ans.members[memberID] = member
	return nil
}

func (ans *LibService) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range ans.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (ans *LibService) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := ans.members[memberID]
	if !ok {
		return nil
	}
	return member.BorrowedBooks
}

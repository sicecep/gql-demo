package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/sicecep/gql-demo/graph/model"
)

// CreateBook is the resolver for the CreateBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	book, err := r.BookRepository.CreateBook(&input)
	bookCreated := &model.Book{
		Author:    book.Author,
		Publisher: book.Publisher,
		Title:     book.Title,
		ID:        book.ID,
	}
	if err != nil {
		return nil, err
	}
	return bookCreated, nil
}

// DeleteBook is the resolver for the DeleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	err := r.BookRepository.DeleteBook(id)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted"
	return successMessage, nil
}

// UpdateBook is the resolver for the UpdateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BookInput) (string, error) {
	err := r.BookRepository.UpdateBook(&input, id)
	if err != nil {
		return "nil", err
	}
	successMessage := "successfully updated"

	return successMessage, nil
}

// GetAllBooks is the resolver for the GetAllBooks field.
func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := r.BookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

// GetOneBook is the resolver for the GetOneBook field.
func (r *queryResolver) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
	book, err := r.BookRepository.GetOneBook(id)
	selectedBook := &model.Book{
		ID:        book.ID,
		Author:    book.Author,
		Publisher: book.Publisher,
		Title:     book.Title,
	}
	if err != nil {
		return nil, err
	}
	return selectedBook, nil
}

// GetOneBookByAuthor is the resolver for the GetOneBookByAuthor field.
func (r *queryResolver) GetOneBookByAuthor(ctx context.Context, author string) (*model.Book, error) {
	book, err := r.BookRepository.GetOneBookByAuthor(author)
	selectedBook := &model.Book{
		ID:        book.ID,
		Author:    book.Author,
		Publisher: book.Publisher,
		Title:     book.Title,
	}
	if err != nil {
		return nil, err
	}
	return selectedBook, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

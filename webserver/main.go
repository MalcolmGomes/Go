package main

import (
    "context"
    "fmt"
    "net/http"
    "os"

    "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
)

// Define the Book struct
type Book struct {
    ID     string
    Title  string
    Author string
}

// Initialize sample data
var books = []Book{
    {ID: "1", Title: "1984", Author: "George Orwell"},
    {ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
}

// Define the QueryResolver
type QueryResolver struct{}

func (r *QueryResolver) Books(ctx context.Context) ([]*BookResolver, error) {
    var res []*BookResolver
    for _, book := range books {
        res = append(res, &BookResolver{book: book})
    }
    return res, nil
}

func (r *QueryResolver) Book(ctx context.Context, args struct{ ID string }) (*BookResolver, error) {
    for _, book := range books {
        if book.ID == args.ID {
            return &BookResolver{book: book}, nil
        }
    }
    return nil, fmt.Errorf("book not found")
}

// Define the MutationResolver
type MutationResolver struct{}

func (r *MutationResolver) AddBook(ctx context.Context, args struct {
    ID     string
    Title  string
    Author string
}) (*BookResolver, error) {
    book := Book{
        ID:     args.ID,
        Title:  args.Title,
        Author: args.Author,
    }
    books = append(books, book)
    return &BookResolver{book: book}, nil
}

// Define the BookResolver
type BookResolver struct {
    book Book
}

func (r *BookResolver) ID() string {
    return r.book.ID
}

func (r *BookResolver) Title() string {
    return r.book.Title
}

func (r *BookResolver) Author() string {
    return r.book.Author
}

// LoadSchema loads the GraphQL schema from a file
func LoadSchema(filename string) (*graphql.Schema, error) {
    schemaData, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    schema := string(schemaData)
    return graphql.MustParseSchema(schema, &RootResolver{}), nil
}

// RootResolver combines QueryResolver and MutationResolver
type RootResolver struct {
    QueryResolver
    MutationResolver
}

func main() {
    schema, err := LoadSchema("schema.graphql")
    if err != nil {
        panic(err)
    }

    http.Handle("/query", &relay.Handler{Schema: schema})
    fmt.Println("Server started at http://localhost:8000")
    http.ListenAndServe(":8000", nil)
}

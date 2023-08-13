# Go CRUD API

### This is a simple Go CRUD (Create, Read, Update, Delete) API using the Gorilla Mux router and Google's UUID package.

### Getting Started

These instructions will help you set up and run the Go CRUD API on your local machine for development and testing purposes.

### Installation

Clone the repository to your local machine:

`git clone https://github.com/vladsw764/CRUD-GoLang-In-Memory.git`

### Navigate to the project directory:

`cd CRUD-GoLang-In-Memory`

### Install the required dependencies using Go modules:

`go mod download`

## Usage

### Run the API server:

`go run main.go`

This will start the API server at http://localhost:8080.

### Interact with the API:

You can use tools like curl, Postman, or even a web browser to interact with the API endpoints.

### Get all books:

GET http://localhost:8080/books

### Get a specific book:

GET http://localhost:8080/books/{id}

### Create a new book:

POST http://localhost:8080/books Body: JSON representation of a book

### Update a book:

PUT http://localhost:8080/books/{id} Body: JSON representation of updated book

### Delete a book:

DELETE http://localhost:8080/books/{id}

## API Endpoints

GET /books - Retrieve a list of all books.

GET /books/{id} - Retrieve details of a specific book.

POST /books - Create a new book.

PUT /books/{id} - Update details of a specific book.

DELETE /books/{id} - Delete a specific book.

## Structure

The main.go file contains the main server code and route definitions.
The Book and Client structs define the data structures for the books and clients.
The API uses UUIDs for generating unique IDs for books.
The Gorilla Mux router is used for handling API routes.

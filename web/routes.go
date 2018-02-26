package web

import "net/http"

const (
	authorBaseUrl   = "/authors"
	authorByUuidUrl = authorBaseUrl + "/{uuid}"
	booksBaseUrl    = "/books"
	bookByUuidUrl   = booksBaseUrl + "/{uuid}"
)

type Route struct {
	//method type (GET, POST, PUT, DELETE, etc)
	Method string

	//the method path
	Pattern string

	//the method that the endpoint should call
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//book_handlers
	Route{
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	Route{
		Method:      "GET",
		Pattern:     booksBaseUrl,
		HandlerFunc: GetAllBooks,
	},
	Route{
		Method:      "POST",
		Pattern:     booksBaseUrl,
		HandlerFunc: AddBook,
	},
	Route{
		Method:      "GET",
		Pattern:     bookByUuidUrl,
		HandlerFunc: GetBookByUUID,
	},
	Route{
		Method:      "DELETE",
		Pattern:     bookByUuidUrl,
		HandlerFunc: DeleteBookByUUID,
	},
	Route{
		Method:      "PUT",
		Pattern:     bookByUuidUrl,
		HandlerFunc: UpdateBook,
	},

	//author_handlers
	Route{
		Method:      "GET",
		Pattern:     authorBaseUrl,
		HandlerFunc: GetAllAuthors,
	},
	Route{
		Method:      "POST",
		Pattern:     authorBaseUrl,
		HandlerFunc: AddAuthor,
	},
	Route{
		Method:      "GET",
		Pattern:     authorByUuidUrl,
		HandlerFunc: GetAuthorByUUID,
	},
	Route{
		Method:      "DELETE",
		Pattern:     authorByUuidUrl,
		HandlerFunc: DeleteAuthorByUUID,
	},
	Route{
		Method:      "PUT",
		Pattern:     authorByUuidUrl,
		HandlerFunc: UpdateAuthor,
	},
}

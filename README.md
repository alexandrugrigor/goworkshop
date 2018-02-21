# golang workshop

# Part 1 - Write REST services with data loaded from files
### 1. hello world in GO
1. create main directory
1. create main.go file
1. declare package main
1. declare main function
1. import fmt package
1. make a call to fmt.Println and print "Hello world, my name is <your_name_here>"
### 2. defining the data model
1. create datamodel.go file
1. create AuthorDto struct with the following fields:
 - UUID string
 - FirstName string
 - LastName string
 - Birthday string
 - Death string
1. create BookDto struct with the following fields:
- UUID string
- Title string
- NoPages int
- ReleaseDate string
- Author AuthorDto
4. create a slice of AuthorDto which will hold all authors in the system
5. create a slice of BookDto which will hold all books in the system
### 3. add JSON marshalling
#### AuthorDto JSON mappings:
* map **AuthorDto.UUID** field to **uuid** field in the resulting JSON
* map **AuthorDto.FirstName** field to **firstName** field in the resulting JSON
* map **AuthorDto.LastName** field to **lastName** field in the resulting JSON
* map **AuthorDto.Birthday** field to **birthday** field in the resulting JSON
* map **AuthorDto.Death** field to **death** field in the resulting JSON
#### BookDto JSON mappings
* map **BookDto.UUID** field to **uuid** field in the resulting JSON
* map **BookDto.Title** field to **title** field in the resulting JSON
* map **BookDto.NoPages** field to **noPages** field in the resulting JSON
* map **BookDto.ReleaseDate** field to **releaseDate** field in the resulting JSON
* map **BookDto.Author** field to **author** field in the resulting JSON
### 4. read sample data from files
1. create model package
1. move datamodel.go in model folder
1. create importer package
1. create authors.json file in importer folder
1. add an array of authors with sample data in JSON format
1. create books.json file in importer folder
1. add an array of books with sample data in JSON format
1. create dataImporter.go file in importer folder
1. import encoding/json package
1. import io/ioutil
1. create a function named ImportAuthors that reads authors.json file and return a slice of authors
1. create a function named ImportBooks that reads the books.json file and returns a slice of books
1. print all authors to the STDOUT
1. print all books to the STDOUT
### 5. start a web server, listening on a configured port
### 6. writing a simple REST endpoint
### 7. implement all REST endpoints with data loaded from files

# Part 2 - Add unit tests for all REST endpoints and create the persistence layer
### 8. write unit tests for all REST endpoints
### 9. add persistence mappings
### 10. add persistence services to retrieve data from db
### 11. switch from loading data from db instead of files

#Part 3 - Add error handling and logging, improve performance with Go routines
### 12. add error handling
### 13. add logging support
### 14. create a configuration service for the application
### 15. improve performance by handling requests async using Go routines

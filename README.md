### Intro
This is a basic backend API for Esquina, a fictional, minimalist app that lets you 
showcase your business to customers in the most glorious of ways.

Alternatively, it is a prototype API to figure out the pros and cons of 
go-swagger and swaggo - The two major packages for working with OpenAPI in Golang. 
### Core Dependencies
- [go-chi/chi](github.com/go-chi/chi) (for API routing).
- [rs/cors](github.com/rs/cors) (for Handling CORS).
- [goswagger](https://goswagger.io) (for code-first documentation).
- [swaggo/swag](github.com/swaggo/swag) (go-swagger alternative because, choices). 
### Requirements
Go - v1.13
### App Setup
* Clone this repository
```bash
$ git clone https://github.com/idoqo/esquina
```
* Enter the project directory with:
```
$ cd esquina
```
* Execute the `main.go` file with:
```
$ go run main.go
```
### Tests
Unit and integration tests are written using the testing library built into go.
Run them with:
```
go test
```
### Documentation
The API documentation is an Open API description powered by go-swagger. 
You can view it at: http://abeg-rest.haq
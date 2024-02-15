# Simple example authentication using JWT in Golang Gin-gonic framework and GORM

## Installation guide

- Clone this project and move to root of the project directory
- Create a `.env` file, copy all variables from `.env.example` file and fill all missing variables
- Open `Mysql`, create database named in `.env` file
- Run `go mod tidy` to install all dependencies
- Run `go run server.go` to start server

## API List

- Header for authentication:
    ```js 
        Authorization: ${token}
    ```
| Endpoint | Method | Auth ? | Payload | Description |
| -------- | -------- | -------- | -------- | -------- |
| **/api/auth/login** | POST | No | *email*, *password* | Login api, return user information and token |
| **/api/auth/register** | POST | No | *email*, *password*, *name* | Register api, return registered user information and token |
| **/api/auth/user** | POST | Yes |  | Get current user information api |
| **/api/books** | GET | Yes |  | Get all books api |
| **/api/books/:id** | GET | Yes |  | Get one book information by id api |
| **/api/books** | POST | Yes | *title*, *page* | Create new book api |
| **/api/books/:id** | PUT | Yes | *title*, *page* | Update book information api, only book creator can do |
| **/api/books/:id** | DELETE | Yes |  | Delete one book by id api, only book creator can do |
| **/api/books** | DELETE | Yes |  | Delete all books of current user |



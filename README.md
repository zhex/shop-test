# ShopX Go Server

This is a simple REST API server built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/) using a SQLite database.

## Features
- List products (`GET /products`)
- Create a product (`POST /products`)

## Prerequisites
- [Go 1.23+](https://golang.org/dl/)

## Installation & Running

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd shop-test
   ```

2. **Download dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   ```bash
   go run *.go
   ```
   The server will start on `http://localhost:8080` by default.


## API Endpoints

- `GET /products` - Returns a list of products in JSON format
- `POST /products` - Creates a new product. Accepts JSON body:


## API Testing with request.http

You can use the `request.http` file (compatible with VS Code REST Client or JetBrains HTTP Client) to quickly test the API endpoints.

To use:
1. Open `request.http` in your editor.
2. Click the "Send Request" button above each request (or use your editor's HTTP client features).

## Database
- Uses SQLite (`products.db` file will be created automatically)
- Seeds with sample products if the database is empty
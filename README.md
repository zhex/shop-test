# ShopX Go Server

This is a simple REST API server built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/) using a SQLite database.

## Features
- List products (`GET /products`)

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

## Database
- Uses SQLite (`products.db` file will be created automatically)
- Seeds with sample products if the database is empty
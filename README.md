
# FastAPI & Go Digital Library

This project is a RESTful Digital Library API built with FastAPI (Python) and Gin (Go), created for the AMI Assignment. It demonstrates core backend concepts including CRUD operations, request validation, proper HTTP error handling, middleware logging, and background task processing using an in-memory data store.

## Features
- Add, view, update, and delete books
- Strict request validation (Python: Pydantic, Go: custom logic)
- Proper HTTP status codes and error handling
- Middleware logging (User-Agent, request time)
- Background task processing (Go)
- In-memory data store (no database required)

---

## Python FastAPI Backend

### Prerequisites
- Python 3.8+
- pip

### Installation
1. Clone the repository:
	 ```bash
	 git clone https://github.com/msaswata15/fastapi-digital-library.git
	 cd fastapi-digital-library/backend
	 ```
2. Install dependencies:
	 ```bash
	 pip install -r ../requirements.txt
	 ```

### Running the Application
Start the FastAPI server with Uvicorn:
```bash
uvicorn main:app --reload
```
The API will be available at [http://127.0.0.1:8000](http://127.0.0.1:8000)

### API Documentation
Interactive docs are available at:
- Swagger UI: [http://127.0.0.1:8000/docs](http://127.0.0.1:8000/docs)
- ReDoc: [http://127.0.0.1:8000/redoc](http://127.0.0.1:8000/redoc)

---

## Go Gin Backend

### Prerequisites
- Go 1.21+

### Installation
1. Clone the repository:
	 ```bash
	 git clone https://github.com/msaswata15/fastapi-digital-library.git
	 cd fastapi-digital-library/go-backend
	 ```
2. Download dependencies:
	 ```bash
	 go mod tidy
	 ```

### Running the Application
Start the Go Gin server:
```bash
go run ./cmd/server
```
The API will be available at [http://localhost:8080](http://localhost:8080)

### API Endpoints

- `POST /items` — Add a new book
- `GET /items` — Get all books
- `GET /items/:id` — Get a book by ID
- `PUT /items/:id` — Update a book (ID in path and body must match)
- `DELETE /items/:id` — Delete a book

#### Example Request (POST /items)
```json
{
	"id": 1,
	"title": "Go Programming",
	"author": "John Doe",
	"year": 2025,
	"isbn": "1234567890"
}
```

#### Example Response
```json
{
	"id": 1,
	"title": "Go Programming",
	"author": "John Doe",
	"year": 2025,
	"isbn": "1234567890"
}
```

#### Error Example (Duplicate ID/ISBN)
```json
{
	"error": "conflict"
}
```

#### Delete Example
```json
{
	"message": "Book removed successfully"
}
```

---

## License
This project is licensed under the MIT License.

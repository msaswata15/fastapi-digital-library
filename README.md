# FastAPI Digital Library

This project is a RESTful Digital Library API built with FastAPI, created for the AMI FastAPI Assignment. It demonstrates core backend concepts including CRUD operations, auto-increment IDs, request validation, and proper HTTP error handling using an in-memory data store.

## Features
- Add, view, update, and delete books
- Auto-incrementing book IDs
- Request validation with Pydantic
- Proper HTTP status codes and error handling
- In-memory data store (no database required)

## Getting Started

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

## Example API Endpoints

### Add a Book
```
POST /books
{
	"title": "Book Title",
	"author": "Author Name",
	"year": 2024
}
```

### Get All Books
```
GET /books
```

### Get a Book by ID
```
GET /books/{book_id}
```

### Update a Book
```
PUT /books/{book_id}
{
	"title": "New Title",
	"author": "New Author",
	"year": 2025
}
```

### Delete a Book
```
DELETE /books/{book_id}
```

## License
This project is licensed under the MIT License.

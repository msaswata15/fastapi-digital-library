from fastapi import FastAPI,HTTPException,Request
import time
from pydantic import BaseModel,Field,field_validator
app=FastAPI(
    title="FastAPI Digital Library",
    description="A simple digital library system to manage books using REST APIs",
    version="1.0.0"
)
@app.middleware("http")
async def logging_middleware(request: Request, call_next):
    start_time = time.time()
    user_agent = request.headers.get("user-agent", "Unknown")
    print(f"[LOG] Request received from: {user_agent}")
    response = await call_next(request)
    process_time = time.time() - start_time
    response.headers["X-Process-Time"] = str(process_time)
    return response

class Book(BaseModel):
    id :int
    title:str =Field(..., min_length=1) 
    author:str
    year:int=Field(ge=1000,le=2026,default=2026)
    isbn:str 
    

    @field_validator('isbn')
    @classmethod
    def isbn_validator(cls, v: str):
            if len(v) not in (10, 13):
                raise ValueError("ISBN must be exactly 10 or 13 digits")
            if not v.isdigit():
                raise ValueError("ISBN must contain only digits")
            return v
books=[]
@app.post("/items",tags=["Library"],
    summary="Add a new book",
    description="Creates a new book entry in the digital library after validating ID and ISBN uniqueness."
)
def createlib(book:Book):
    for bookt in books:
        if bookt.id==book.id or bookt.isbn==book.isbn:
            raise HTTPException(status_code=400,
                detail=f"Book with id {book.id} or book with isbn {book.isbn} already exists")
    books.append(book)
    return book

@app.get("/items/{item_id}",tags=["Library"],
    summary="Get book by ID",
    description="Fetches a single book from the library using its unique ID."
)
def get_byid(item_id:int):
    for book in books:
        if book.id==item_id:
            return book
          
    raise HTTPException(status_code=404,detail="Book id not present please enter a valid book id.")

@app.get("/items",tags=["Library"],
    summary="Get all books",
    description="Retrieves the complete list of books available in the digital library."
)
def get_all():
     return books

@app.delete("/items/{item_id}",tags=["Library"],
    summary="Delete a book",
    description="Removes a book from the digital library using its ID."
)
def delete_book(item_id:int):
    for book in books:
        if book.id==item_id:
            
            books.remove(book)
            return {"message": "Book removed successfully"}

    
    raise HTTPException(status_code=404,detail="Book not found")

    
@app.put("/items/{book_id}", response_model=Book,tags=['Library'],
    summary="Update book details",
    description="Updates all details of an existing book. Book ID in path and body must match."
)
def update_book(book_id: int, updated_book: Book):
    if updated_book.id != book_id:
        raise HTTPException(
            status_code=400,
            detail="Book ID in path and body must match"
        )

    for index, book in enumerate(books):
        if book.id == book_id:
            books[index] = updated_book
            return updated_book

    raise HTTPException(
        status_code=404,
        detail=f"Book with id {book_id} not found"
    ) 
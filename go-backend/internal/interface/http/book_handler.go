package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"fastapi-digital-library/go-backend/internal/domain"
	"fastapi-digital-library/go-backend/internal/usecase"
)

type BookHandler struct {
	service *usecase.BookService
	queue   *TaskQueue
}

func NewBookHandler(s *usecase.BookService, q *TaskQueue) *BookHandler {
	return &BookHandler{service: s, queue: q}
}

func (h *BookHandler) Register(r *gin.Engine) {
	grp := r.Group("/items")
	grp.POST("", h.create)
	grp.GET("", h.getAll)
	grp.GET(":id", h.getByID)
	grp.DELETE(":id", h.delete)
	grp.PUT(":id", h.update)
}

func (h *BookHandler) create(c *gin.Context) {
	var b domain.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		writeError(c, http.StatusBadRequest, "invalid payload")
		return
	}
	created, err := h.service.Create(b)
	if err != nil {
		status := mapError(err)
		writeError(c, status, err.Error())
		return
	}
	if h.queue != nil {
		h.queue.Enqueue(Task{Type: "index_book", Payload: created})
	}
	c.JSON(http.StatusOK, created)
}

func (h *BookHandler) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		writeError(c, http.StatusBadRequest, "invalid id")
		return
	}
	book, err := h.service.GetByID(id)
	if err != nil {
		status := mapError(err)
		writeError(c, status, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) getAll(c *gin.Context) {
	books, err := h.service.GetAll()
	if err != nil {
		writeError(c, http.StatusInternalServerError, "failed to fetch books")
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		writeError(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.service.Delete(id); err != nil {
		status := mapError(err)
		writeError(c, status, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
}

func (h *BookHandler) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		writeError(c, http.StatusBadRequest, "invalid id")
		return
	}
	var b domain.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		writeError(c, http.StatusBadRequest, "invalid payload")
		return
	}
	if b.ID != id {
		writeError(c, http.StatusBadRequest, "Book ID in path and body must match")
		return
	}
	updated, err := h.service.Update(b)
	if err != nil {
		status := mapError(err)
		writeError(c, status, err.Error())
		return
	}
	c.JSON(http.StatusOK, updated)
}

func mapError(err error) int {
	switch err {
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict, domain.ErrValidation:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func writeError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"error": msg})
}

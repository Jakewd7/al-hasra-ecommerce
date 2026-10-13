package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new product category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        input  body      CreateCategoryRequest  true  "Category Data"
// @Success      201    {object}  Category
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/categories [post]
func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetAll godoc
// @Summary      Get all categories
// @Description  Get list of all product categories
// @Tags         categories
// @Produce      json
// @Success      200  {array}   Category
// @Failure      500  {object}  map[string]string
// @Router       /api/categories [get]
func (h *Handler) GetAll(c *gin.Context) {
	results, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// GetByID godoc
// @Summary      Get category by ID
// @Description  Get details of a single category by its UUID
// @Tags         categories
// @Produce      json
// @Param        id   path      string  true  "Category UUID"
// @Success      200  {object}  Category
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/categories/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	result, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Update godoc
// @Summary      Update a category
// @Description  Update category details by UUID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id     path      string                 true  "Category UUID"
// @Param        input  body      UpdateCategoryRequest  true  "Updated Category Data"
// @Success      200    {object}  Category
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/categories/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.service.Update(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Delete godoc
// @Summary      Delete a category
// @Description  Delete category by UUID
// @Tags         categories
// @Produce      json
// @Param        id   path      string  true  "Category UUID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/categories/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

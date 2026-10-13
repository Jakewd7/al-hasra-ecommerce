package roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// GetRoles godoc
// @Summary Get all roles
// @Description Get list of roles including their permissions
// @Tags Roles
// @Produce json
// @Success 200 {array} RoleResponse
// @Router /roles [get]
func (h *Handler) GetRoles(c *gin.Context) {
	roles, err := h.service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// CreateRole godoc
// @Summary Create new role
// @Description Create a new role and assign permissions
// @Tags Roles
// @Accept json
// @Produce json
// @Param request body CreateRoleRequest true "Create Role Request"
// @Success 201 {object} RoleResponse
// @Failure 400 {object} map[string]string
// @Router /roles [post]
func (h *Handler) CreateRole(c *gin.Context) {
	var input CreateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.service.CreateRole(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

// UpdateRole godoc
// @Summary Update role
// @Description Update role details and permissions
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param request body UpdateRoleRequest true "Update Role Request"
// @Success 200 {object} RoleResponse
// @Failure 400 {object} map[string]string
// @Router /roles/{id} [put]
func (h *Handler) UpdateRole(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var input UpdateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.service.UpdateRole(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

// DeleteRole godoc
// @Summary Delete role
// @Description Delete role by ID
// @Tags Roles
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /roles/{id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := h.service.DeleteRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// GetPermissions godoc
// @Summary Get all permissions
// @Description Get list of all available permissions (useful for frontend selection)
// @Tags Roles
// @Produce json
// @Success 200 {array} PermissionResponse
// @Router /permissions [get]
func (h *Handler) GetPermissions(c *gin.Context) {
	perms, err := h.service.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, perms)
}

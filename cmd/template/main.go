package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ModuleData struct {
	Name         string
	InstanceName string
	PackageName  string
	RoutePath    string
	ImportPath   string
}

func main() {
	var input string
	fmt.Print("Enter module name (e.g., transaction-item): ")
	fmt.Scanln(&input)

	if input == "" {
		fmt.Println("Error: Module name cannot be empty")
		return
	}

	data := processName(input)
	data.ImportPath = "github.com/jakewd7"

	modulePath := filepath.Join("internal", "modules", data.PackageName)
	if err := os.MkdirAll(modulePath, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	templates := map[string]string{
		"model.go":      modelTemplate,
		"repository.go": repoTemplate,
		"service.go":    serviceTemplate,
		"handler.go":    handlerTemplate,
		"routes.go":     routeTemplate,
	}

	for fileName, content := range templates {
		filePath := filepath.Join(modulePath, fileName)
		createFileFromTemplate(filePath, content, data)
	}

	updateGlobalRoutes(data)

	fmt.Printf("\n🚀 Module '%s' generated successfully!\n", data.Name)
	fmt.Printf("📍 Path: %s\n", modulePath)
	fmt.Println("✅ Global routes updated.")
}

func processName(input string) ModuleData {
	parts := strings.Split(input, "-")

	var pascalParts []string
	for _, part := range parts {
		pascalParts = append(pascalParts, strings.Title(part))
	}

	pascalName := strings.Join(pascalParts, "")
	camelName := strings.ToLower(pascalParts[0]) + strings.Join(pascalParts[1:], "")
	rawName := strings.ToLower(strings.Join(parts, ""))
	packageName := pluralize(rawName)
	routePath := pluralize(input)

	return ModuleData{
		Name:         pascalName,
		InstanceName: camelName,
		PackageName:  packageName,
		RoutePath:    routePath,
	}
}

func pluralize(word string) string {
	if strings.HasSuffix(word, "s") ||
		strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "z") ||
		strings.HasSuffix(word, "ch") ||
		strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	if strings.HasSuffix(word, "y") && len(word) > 1 {
		return word[:len(word)-1] + "ies"
	}

	return word + "s"
}

func updateGlobalRoutes(data ModuleData) {
	routesPath := filepath.Join("internal", "routes", "routes.go")
	content, err := os.ReadFile(routesPath)
	if err != nil {
		fmt.Printf("Warning: Could not read routes.go: %v\n", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	importStr := fmt.Sprintf("\t\"%s/internal/modules/%s\"", data.ImportPath, data.PackageName)
	registerStr := fmt.Sprintf("\t%s.RegisterRoutes(protectedGroup, db)", data.PackageName)

	for _, line := range lines {
		newLines = append(newLines, line)
		if strings.Contains(line, "// [AUTO-IMPORT]") {
			newLines = append(newLines, importStr)
		}
		if strings.Contains(line, "// [AUTO-REGISTER]") {
			newLines = append(newLines, registerStr)
		}
	}

	output := strings.Join(newLines, "\n")
	err = os.WriteFile(routesPath, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Warning: Could not update routes.go: %v\n", err)
	}
}

func createFileFromTemplate(path string, tempStr string, data ModuleData) {
	processedTemp := strings.ReplaceAll(tempStr, "[[BT]]", "`")

	tmpl, err := template.New("gen").Parse(processedTemp)
	if err != nil {
		fmt.Printf("Error parsing template for %s: %v\n", path, err)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", path, err)
		return
	}
	defer file.Close()

	tmpl.Execute(file, data)
}

const modelTemplate = `package {{.PackageName}}

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{.Name}} struct {
	ID          uuid.UUID [[BT]]gorm:"type:uuid;primaryKey"[[BT]]
	Name        string    [[BT]]gorm:"type:varchar(255);not null"[[BT]]
	Description string    [[BT]]gorm:"type:text"[[BT]]
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
}

func (m *{{.Name}}) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

type Create{{.Name}}Request struct {
	Name        string [[BT]]json:"name" binding:"required"[[BT]]
	Description string [[BT]]json:"description"[[BT]]
}

type Update{{.Name}}Request struct {
	Name        string [[BT]]json:"name"[[BT]]
	Description string [[BT]]json:"description"[[BT]]
}
`

const repoTemplate = `package {{.PackageName}}

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Repository interface {
	Create(data *{{.Name}}) error
	FindAll() ([]{{.Name}}, error)
	FindByID(id uuid.UUID) (*{{.Name}}, error)
	Update(data *{{.Name}}) error
	Delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(data *{{.Name}}) error {
	return r.db.Create(data).Error
}

func (r *repository) FindAll() ([]{{.Name}}, error) {
	var results []{{.Name}}
	err := r.db.Find(&results).Error
	return results, err
}

func (r *repository) FindByID(id uuid.UUID) (*{{.Name}}, error) {
	var result {{.Name}}
	err := r.db.First(&result, "id = ?", id).Error
	return &result, err
}

func (r *repository) Update(data *{{.Name}}) error {
	return r.db.Save(data).Error
}

func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&{{.Name}}{}, "id = ?", id).Error
}
`

const serviceTemplate = `package {{.PackageName}}

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(req Create{{.Name}}Request) (*{{.Name}}, error) {
	data := &{{.Name}}{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.repo.Create(data)
	return data, err
}

func (s *Service) GetAll() ([]{{.Name}}, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*{{.Name}}, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id uuid.UUID, req Update{{.Name}}Request) (*{{.Name}}, error) {
	data, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	
	if req.Name != "" {
		data.Name = req.Name
	}
	data.Description = req.Description

	err = s.repo.Update(data)
	return data, err
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
`

const handlerTemplate = `package {{.PackageName}}

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Create{{.Name}}(c *gin.Context) {
	var req Create{{.Name}}Request
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

func (h *Handler) GetAll(c *gin.Context) {
	results, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	result, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var req Update{{.Name}}Request
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

func (h *Handler) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
`

const routeTemplate = `package {{.PackageName}}

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	{{.InstanceName}} := r.Group("/{{.RoutePath}}")
	{
		{{.InstanceName}}.POST("", handler.Create{{.Name}})
		{{.InstanceName}}.GET("", handler.GetAll)
		{{.InstanceName}}.GET("/:id", handler.GetByID)
		{{.InstanceName}}.PUT("/:id", handler.Update)
		{{.InstanceName}}.DELETE("/:id", handler.Delete)
	}
}
`

package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"kasir-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCategoryService is a mock of CategoryService
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetAll() ([]models.Category, error) {
	args := m.Called()
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) GetByID(id int) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryService) Create(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryService) Update(category *models.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllCategories(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	expectedCategories := []models.Category{
		{ID: 1, Name: "Category 1", Description: "Description 1"},
		{ID: 2, Name: "Category 2", Description: "Description 2"},
	}

	mockService.On("GetAll").Return(expectedCategories, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/kategori", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetAll(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var categories []models.Category
	err = json.Unmarshal(rr.Body.Bytes(), &categories)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(categories))
	assert.Equal(t, expectedCategories[0].Name, categories[0].Name)

	mockService.AssertExpectations(t)
}

func TestGetAllCategories_ServiceError(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	mockService.On("GetAll").Return([]models.Category{}, errors.New("database error"))

	req, err := http.NewRequest(http.MethodGet, "/api/kategori", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetAll(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestCreateCategory(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	newCategory := models.Category{
		Name:        "New Category",
		Description: "New Description",
	}

	mockService.On("Create", mock.AnythingOfType("*models.Category")).Return(nil)

	body, _ := json.Marshal(newCategory)
	req, err := http.NewRequest(http.MethodPost, "/api/kategori", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response models.Category
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newCategory.Name, response.Name)

	mockService.AssertExpectations(t)
}

func TestCreateCategory_InvalidJSON(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	req, err := http.NewRequest(http.MethodPost, "/api/kategori", bytes.NewBufferString("invalid json"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateCategory_ServiceError(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	newCategory := models.Category{
		Name:        "New Category",
		Description: "New Description",
	}

	mockService.On("Create", mock.AnythingOfType("*models.Category")).Return(errors.New("validation error"))

	body, _ := json.Marshal(newCategory)
	req, err := http.NewRequest(http.MethodPost, "/api/kategori", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetCategoryByID(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	expectedCategory := &models.Category{
		ID:          1,
		Name:        "Category 1",
		Description: "Description 1",
	}

	mockService.On("GetByID", 1).Return(expectedCategory, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/kategori/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var category models.Category
	err = json.Unmarshal(rr.Body.Bytes(), &category)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory.Name, category.Name)

	mockService.AssertExpectations(t)
}

func TestGetCategoryByID_InvalidID(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	req, err := http.NewRequest(http.MethodGet, "/api/kategori/invalid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetCategoryByID_NotFound(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	mockService.On("GetByID", 999).Return((*models.Category)(nil), errors.New("category not found"))

	req, err := http.NewRequest(http.MethodGet, "/api/kategori/999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	updatedCategory := models.Category{
		ID:          1,
		Name:        "Updated Category",
		Description: "Updated Description",
	}

	mockService.On("Update", mock.AnythingOfType("*models.Category")).Return(nil)

	body, _ := json.Marshal(updatedCategory)
	req, err := http.NewRequest(http.MethodPut, "/api/kategori/1", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateCategory_InvalidID(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	updatedCategory := models.Category{
		Name:        "Updated Category",
		Description: "Updated Description",
	}

	body, _ := json.Marshal(updatedCategory)
	req, err := http.NewRequest(http.MethodPut, "/api/kategori/invalid", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateCategory_InvalidJSON(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	req, err := http.NewRequest(http.MethodPut, "/api/kategori/1", bytes.NewBufferString("invalid json"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateCategory_ServiceError(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	updatedCategory := models.Category{
		ID:          1,
		Name:        "Updated Category",
		Description: "Updated Description",
	}

	mockService.On("Update", mock.AnythingOfType("*models.Category")).Return(errors.New("update failed"))

	body, _ := json.Marshal(updatedCategory)
	req, err := http.NewRequest(http.MethodPut, "/api/kategori/1", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	mockService.On("Delete", 1).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, "/api/kategori/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "category deleted successfully", response["message"])

	mockService.AssertExpectations(t)
}

func TestDeleteCategory_InvalidID(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	req, err := http.NewRequest(http.MethodDelete, "/api/kategori/invalid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteCategory_ServiceError(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	mockService.On("Delete", 999).Return(errors.New("category not found"))

	req, err := http.NewRequest(http.MethodDelete, "/api/kategori/999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandleCategories_MethodNotAllowed(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	req, err := http.NewRequest(http.MethodPatch, "/api/kategori", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.HandleCategories(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestHandleCategoryByID_MethodNotAllowed(t *testing.T) {
	mockService := new(MockCategoryService)
	handler := NewCategoryHandler(mockService)

	req, err := http.NewRequest(http.MethodPatch, "/api/kategori/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.HandleCategoryByID(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

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

// MockProductService is a mock of ProductService
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) GetByID(id int) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) Create(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllProducts(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	expectedProducts := []models.Product{
		{ID: 1, Name: "Product 1", Price: 10000, Stock: 50},
		{ID: 2, Name: "Product 2", Price: 20000, Stock: 30},
	}

	mockService.On("GetAll").Return(expectedProducts, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/produk", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetAll(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var products []models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)

	mockService.AssertExpectations(t)
}

func TestGetAllProducts_ServiceError(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	mockService.On("GetAll").Return([]models.Product{}, errors.New("database error"))

	req, err := http.NewRequest(http.MethodGet, "/api/produk", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetAll(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	newProduct := models.Product{
		Name:  "New Product",
		Price: 15000,
		Stock: 100,
	}

	mockService.On("Create", mock.AnythingOfType("*models.Product")).Return(nil)

	body, _ := json.Marshal(newProduct)
	req, err := http.NewRequest(http.MethodPost, "/api/produk", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newProduct.Name, response.Name)

	mockService.AssertExpectations(t)
}

func TestCreateProduct_InvalidJSON(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	req, err := http.NewRequest(http.MethodPost, "/api/produk", bytes.NewBufferString("invalid json"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateProduct_ServiceError(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	newProduct := models.Product{
		Name:  "New Product",
		Price: 15000,
		Stock: 100,
	}

	mockService.On("Create", mock.AnythingOfType("*models.Product")).Return(errors.New("validation error"))

	body, _ := json.Marshal(newProduct)
	req, err := http.NewRequest(http.MethodPost, "/api/produk", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockService.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	expectedProduct := &models.Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10000,
		Stock: 50,
	}

	mockService.On("GetByID", 1).Return(expectedProduct, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/produk/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var product models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &product)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct.Name, product.Name)

	mockService.AssertExpectations(t)
}

func TestGetProductByID_InvalidID(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	req, err := http.NewRequest(http.MethodGet, "/api/produk/invalid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetProductByID_NotFound(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	mockService.On("GetByID", 999).Return((*models.Product)(nil), errors.New("product not found"))

	req, err := http.NewRequest(http.MethodGet, "/api/produk/999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	updatedProduct := models.Product{
		ID:    1,
		Name:  "Updated Product",
		Price: 25000,
		Stock: 75,
	}

	mockService.On("Update", mock.AnythingOfType("*models.Product")).Return(nil)

	body, _ := json.Marshal(updatedProduct)
	req, err := http.NewRequest(http.MethodPut, "/api/produk/1", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateProduct_InvalidID(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	updatedProduct := models.Product{
		Name:  "Updated Product",
		Price: 25000,
		Stock: 75,
	}

	body, _ := json.Marshal(updatedProduct)
	req, err := http.NewRequest(http.MethodPut, "/api/produk/invalid", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateProduct_InvalidJSON(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	req, err := http.NewRequest(http.MethodPut, "/api/produk/1", bytes.NewBufferString("invalid json"))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateProduct_ServiceError(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	updatedProduct := models.Product{
		ID:    1,
		Name:  "Updated Product",
		Price: 25000,
		Stock: 75,
	}

	mockService.On("Update", mock.AnythingOfType("*models.Product")).Return(errors.New("update failed"))

	body, _ := json.Marshal(updatedProduct)
	req, err := http.NewRequest(http.MethodPut, "/api/produk/1", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	mockService.On("Delete", 1).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, "/api/produk/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Product deleted successfully", response["message"])

	mockService.AssertExpectations(t)
}

func TestDeleteProduct_InvalidID(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	req, err := http.NewRequest(http.MethodDelete, "/api/produk/invalid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteProduct_ServiceError(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	mockService.On("Delete", 999).Return(errors.New("product not found"))

	req, err := http.NewRequest(http.MethodDelete, "/api/produk/999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Delete(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandleProducts_MethodNotAllowed(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	req, err := http.NewRequest(http.MethodPatch, "/api/produk", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.HandleProducts(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestHandleProductByID_MethodNotAllowed(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	req, err := http.NewRequest(http.MethodPatch, "/api/produk/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.HandleProductByID(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

package handlers

import (
	"bytes"
	"encoding/json"
	"kasir-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestProductWithCategoryIntegration tests that products return category names
func TestProductWithCategoryIntegration(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	// Test GetByID returns product with category name
	t.Run("GetByID returns product with category name", func(t *testing.T) {
		expectedProduct := &models.Product{
			ID:           1,
			Name:         "Laptop Dell XPS",
			Price:        15000000,
			Stock:        5,
			CategoryID:   1,
			CategoryName: "Electronics",
		}

		mockService.On("GetByID", 1).Return(expectedProduct, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/produk/1", nil)
		rr := httptest.NewRecorder()

		handler.GetByID(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var product models.Product
		json.Unmarshal(rr.Body.Bytes(), &product)

		assert.Equal(t, "Laptop Dell XPS", product.Name)
		assert.Equal(t, 1, product.CategoryID)
		assert.Equal(t, "Electronics", product.CategoryName)
	})

	// Test GetAll returns products with category names
	t.Run("GetAll returns products with category names", func(t *testing.T) {
		expectedProducts := []models.Product{
			{
				ID:           1,
				Name:         "Laptop",
				Price:        15000000,
				Stock:        5,
				CategoryID:   1,
				CategoryName: "Electronics",
			},
			{
				ID:           2,
				Name:         "Chair",
				Price:        500000,
				Stock:        20,
				CategoryID:   2,
				CategoryName: "Furniture",
			},
			{
				ID:           3,
				Name:         "Mystery Item",
				Price:        100000,
				Stock:        10,
				CategoryID:   0,
				CategoryName: "", // No category assigned
			},
		}

		mockService.On("GetAll").Return(expectedProducts, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/produk", nil)
		rr := httptest.NewRecorder()

		handler.GetAll(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var products []models.Product
		json.Unmarshal(rr.Body.Bytes(), &products)

		assert.Equal(t, 3, len(products))
		assert.Equal(t, "Electronics", products[0].CategoryName)
		assert.Equal(t, "Furniture", products[1].CategoryName)
		assert.Equal(t, "", products[2].CategoryName) // Product without category
	})

	// Test creating product with category_id
	t.Run("Create product with category_id", func(t *testing.T) {
		newProduct := models.Product{
			Name:       "Mouse Gaming",
			Price:      250000,
			Stock:      50,
			CategoryID: 1,
		}

		mockService.On("Create", &newProduct).Return(nil).Once()

		body, _ := json.Marshal(newProduct)
		req, _ := http.NewRequest(http.MethodPost, "/api/produk", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.Create(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		var createdProduct models.Product
		json.Unmarshal(rr.Body.Bytes(), &createdProduct)

		assert.Equal(t, "Mouse Gaming", createdProduct.Name)
		assert.Equal(t, 1, createdProduct.CategoryID)
	})

	// Test updating product with different category_id
	t.Run("Update product category", func(t *testing.T) {
		updatedProduct := models.Product{
			ID:         1,
			Name:       "Laptop Updated",
			Price:      16000000,
			Stock:      3,
			CategoryID: 2, // Changed from Electronics to Furniture
		}

		mockService.On("Update", &updatedProduct).Return(nil).Once()

		body, _ := json.Marshal(updatedProduct)
		req, _ := http.NewRequest(http.MethodPut, "/api/produk/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler.Update(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	mockService.AssertExpectations(t)
}

// TestProductCategoryJoinScenarios tests various JOIN scenarios
func TestProductCategoryJoinScenarios(t *testing.T) {
	mockService := new(MockProductService)
	handler := NewProductHandler(mockService)

	t.Run("Product with existing category", func(t *testing.T) {
		product := &models.Product{
			ID:           100,
			Name:         "Product A",
			Price:        1000,
			Stock:        10,
			CategoryID:   5,
			CategoryName: "Category A",
		}

		mockService.On("GetByID", 100).Return(product, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/produk/100", nil)
		rr := httptest.NewRecorder()

		handler.GetByID(rr, req)

		var result models.Product
		json.Unmarshal(rr.Body.Bytes(), &result)

		assert.NotEmpty(t, result.CategoryName)
		assert.Equal(t, "Category A", result.CategoryName)
	})

	t.Run("Product without category (LEFT JOIN result)", func(t *testing.T) {
		product := &models.Product{
			ID:           101,
			Name:         "Product B",
			Price:        2000,
			Stock:        5,
			CategoryID:   0,
			CategoryName: "", // Empty because no category assigned
		}

		mockService.On("GetByID", 101).Return(product, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/produk/101", nil)
		rr := httptest.NewRecorder()

		handler.GetByID(rr, req)

		var result models.Product
		json.Unmarshal(rr.Body.Bytes(), &result)

		// Should still work even without category
		assert.Equal(t, "", result.CategoryName)
		assert.Equal(t, "Product B", result.Name)
	})

	mockService.AssertExpectations(t)
}

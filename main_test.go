package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCategories(t *testing.T) {
	// Reset categories for test
	categories = []Category{
		{ID: 1, Name: "Makanan", Description: "Produk makanan"},
		{ID: 2, Name: "Minuman", Description: "Produk minuman"},
	}

	req, err := http.NewRequest("GET", "/api/categories", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getCategories(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"name":"Makanan","description":"Produk makanan"},{"id":2,"name":"Minuman","description":"Produk minuman"}]` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func TestCreateCategory(t *testing.T) {
	// Reset categories
	categories = []Category{}

	category := Category{Name: "Snack", Description: "Produk snack"}
	jsonData, _ := json.Marshal(category)

	req, err := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createCategory(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response Category
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.ID != 1 || response.Name != "Snack" || response.Description != "Produk snack" {
		t.Errorf("handler returned unexpected category: got %v", response)
	}
}

func TestGetCategoryByID(t *testing.T) {
	categories = []Category{
		{ID: 1, Name: "Makanan", Description: "Produk makanan"},
	}

	req, err := http.NewRequest("GET", "/api/categories/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getCategoryByID(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Category
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.ID != 1 || response.Name != "Makanan" {
		t.Errorf("handler returned unexpected category: got %v", response)
	}
}

func TestGetCategoryByIDNotFound(t *testing.T) {
	categories = []Category{}

	req, err := http.NewRequest("GET", "/api/categories/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getCategoryByID(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestUpdateCategory(t *testing.T) {
	categories = []Category{
		{ID: 1, Name: "Makanan", Description: "Produk makanan"},
	}

	updatedCategory := Category{Name: "Food", Description: "Food products"}
	jsonData, _ := json.Marshal(updatedCategory)

	req, err := http.NewRequest("PUT", "/api/categories/1", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateCategory(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Category
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.ID != 1 || response.Name != "Food" || response.Description != "Food products" {
		t.Errorf("handler returned unexpected category: got %v", response)
	}
}

func TestUpdateCategoryNotFound(t *testing.T) {
	categories = []Category{}

	updatedCategory := Category{Name: "Food", Description: "Food products"}
	jsonData, _ := json.Marshal(updatedCategory)

	req, err := http.NewRequest("PUT", "/api/categories/999", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateCategory(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestDeleteCategory(t *testing.T) {
	categories = []Category{
		{ID: 1, Name: "Makanan", Description: "Produk makanan"},
		{ID: 2, Name: "Minuman", Description: "Produk minuman"},
	}

	req, err := http.NewRequest("DELETE", "/api/categories/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteCategory(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if len(categories) != 1 {
		t.Errorf("expected categories length 1, got %v", len(categories))
	}

	if categories[0].ID != 2 {
		t.Errorf("expected remaining category ID 2, got %v", categories[0].ID)
	}
}

func TestDeleteCategoryNotFound(t *testing.T) {
	categories = []Category{}

	req, err := http.NewRequest("DELETE", "/api/categories/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteCategory(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

// Product tests

func TestGetProduk(t *testing.T) {
	// Reset produk for test
	produk = []Produk{
		{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
		{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	}

	req, err := http.NewRequest("GET", "/api/produk", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produk)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"nama":"Indomie Godog","harga":3500,"stok":10},{"id":2,"nama":"Vit 1000ml","harga":3000,"stok":40}]` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func TestCreateProduk(t *testing.T) {
	// Reset produk
	produk = []Produk{}

	produkBaru := Produk{Nama: "Kecap", Harga: 12000, Stok: 20}
	jsonData, _ := json.Marshal(produkBaru)

	req, err := http.NewRequest("POST", "/api/produk", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// baca data dari request
		var produkBaru Produk
		err := json.NewDecoder(r.Body).Decode(&produkBaru)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// masukkin data ke dalam variable produk
		produkBaru.ID = len(produk) + 1
		produk = append(produk, produkBaru)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
		json.NewEncoder(w).Encode(produkBaru)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response Produk
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.ID != 1 || response.Nama != "Kecap" || response.Harga != 12000 || response.Stok != 20 {
		t.Errorf("handler returned unexpected produk: got %v", response)
	}
}

func TestGetProdukByID(t *testing.T) {
	produk = []Produk{
		{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	}

	req, err := http.NewRequest("GET", "/api/produk/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getProdukByID(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Produk
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.ID != 1 || response.Nama != "Indomie Godog" {
		t.Errorf("handler returned unexpected produk: got %v", response)
	}
}

func TestGetProdukByIDNotFound(t *testing.T) {
	produk = []Produk{}

	req, err := http.NewRequest("GET", "/api/produk/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getProdukByID(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestUpdateProduk(t *testing.T) {
	produk = []Produk{
		{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	}

	updatedProduk := Produk{Nama: "Indomie Goreng", Harga: 4000, Stok: 15}
	jsonData, _ := json.Marshal(updatedProduk)

	req, err := http.NewRequest("PUT", "/api/produk/1", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateProduk(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Produk
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.ID != 1 || response.Nama != "Indomie Goreng" || response.Harga != 4000 || response.Stok != 15 {
		t.Errorf("handler returned unexpected produk: got %v", response)
	}
}

func TestUpdateProdukNotFound(t *testing.T) {
	produk = []Produk{}

	updatedProduk := Produk{Nama: "Indomie Goreng", Harga: 4000, Stok: 15}
	jsonData, _ := json.Marshal(updatedProduk)

	req, err := http.NewRequest("PUT", "/api/produk/999", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateProduk(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestDeleteProduk(t *testing.T) {
	produk = []Produk{
		{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
		{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	}

	req, err := http.NewRequest("DELETE", "/api/produk/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteProduk(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if len(produk) != 1 {
		t.Errorf("expected produk length 1, got %v", len(produk))
	}

	if produk[0].ID != 2 {
		t.Errorf("expected remaining produk ID 2, got %v", produk[0].ID)
	}
}

func TestDeleteProdukNotFound(t *testing.T) {
	produk = []Produk{}

	req, err := http.NewRequest("DELETE", "/api/produk/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteProduk(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

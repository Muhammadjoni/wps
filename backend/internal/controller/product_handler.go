package controller

import (
	"backend/internal/model/dto"
	"backend/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	service  service.ProductService
	validate *validator.Validate
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(r.Context(), &req)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Similar handler methods for other endpoints

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/HlufD/products-ms/internal/adapters/left/http/dto"
	"github.com/HlufD/products-ms/internal/core/domain"
	"github.com/HlufD/products-ms/internal/core/usecases"
	"github.com/HlufD/products-ms/shared"
)

type ProductController struct {
	productService *usecases.ProductUseCase
}

func NewProductController(productService *usecases.ProductUseCase) *ProductController {
	return &ProductController{productService}
}

func (pc *ProductController) RegisterRoutes(router chi.Router) {
	router.Route("/api/v1/products", func(r chi.Router) {
		r.Post("/", pc.CreateProduct)
		r.Get("/", pc.GetAllProducts)
		r.Get("/{id}", pc.GetProductByID)
		r.Put("/{id}", pc.UpdateProduct)
		r.Get("/{id}/availability", pc.CheckAvailability)
	})
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product *domain.Product
	var createProductDto dto.CreateProduct

	if err := json.NewDecoder(r.Body).Decode(&createProductDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	// validate it

	if err := shared.Validate(createProductDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// map dto to entity
	product = createProductDto.MapToDomainEntity()

	createdProduct, err := pc.productService.CreateProduct(product)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	shared.RespondWithJSON(w, http.StatusCreated, createdProduct)
}

func (pc *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	println(id)

	product, err := pc.productService.GetProductById(id)

	if product == nil {
		shared.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var product *domain.UpdateProduct
	var updatedProductDto dto.UpdateProduct

	if err := json.NewDecoder(r.Body).Decode(&updatedProductDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	// validate
	if err := shared.Validate(updatedProductDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// map dto to entity
	product = updatedProductDto.MapToDomainEntity()

	updatedProduct, err := pc.productService.UpdateProduct(id, product)

	log.Println(updatedProduct)

	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, updatedProduct)
}

func (pc *ProductController) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	q := r.URL.Query().Get("quantity")

	quantity, err := strconv.Atoi(q)
	if err != nil || quantity <= 0 {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid or missing 'quantity' query parameter")
		return
	}

	available, err := pc.productService.CheckAvailability(id, quantity)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, "Error checking availability")
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, map[string]bool{"available": available})
}

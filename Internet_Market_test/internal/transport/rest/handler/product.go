package handler

import (
	"errors"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/pkg"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	idStr := r.Header.Get("id")
	log.Println(idStr)
	idSeller, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	err = pkg.ReadJSON(w, r, product)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Bad Request"), http.StatusBadRequest)
		return
	}

	idProduct, err := h.service.Product.CreateProduct(product, int32(idSeller))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = idProduct
	payload.Message = "Product Created"
	err = pkg.WriteJSON(w, http.StatusCreated, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) GetProductsSeller(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	idSeller, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	products, err := h.service.Product.GetProductsSeller(int32(idSeller))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = products
	payload.Message = "Products gets"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	product, err := h.service.Product.GetProductByID(int32(id))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = product
	payload.Message = "Product gets"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	idStr := r.Header.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	err = pkg.ReadJSON(w, r, product)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Bad Request"), http.StatusBadRequest)
		return
	}
	err = h.service.Product.UpdateProduct(product, int32(id))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = product
	payload.Message = "Product updated"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	err = h.service.Product.DeleteProduct(int32(id))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Message = "Product deleted"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.Product.GetProducts()
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = products
	payload.Message = "Products gets"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

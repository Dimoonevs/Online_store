package handler

import (
	"errors"
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/pkg"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	seller := &models.Seller{}

	err := pkg.ReadJSON(w, r, seller)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	id, err := h.service.Seller.CreateSeller(seller)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = id
	payload.Message = "Created successfully"
	log.Println(payload)
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) GetSellers(w http.ResponseWriter, r *http.Request) {
	models, err := h.service.Seller.GetSellers()
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = models
	payload.Message = "Retrieved successfully"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) GetSellerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	model, err := h.service.Seller.GetSellerByID(int32(id))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = model
	payload.Message = "Retrieved successfully"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	seller := &models.Seller{}

	err = pkg.ReadJSON(w, r, seller)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.Seller.UpdateSeller(seller, int32(id))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Data = seller
	payload.Message = "Updated successfully"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
func (h *Handler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	idStr := r.Header.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ErrorJSON(w, errors.New("Invalid ID format"), http.StatusBadRequest)
		return
	}
	err = h.service.Seller.DeleteSeller(int32(id))
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	var payload pkg.JsonPayload
	payload.Error = false
	payload.Message = "Deleted successfully"
	err = pkg.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

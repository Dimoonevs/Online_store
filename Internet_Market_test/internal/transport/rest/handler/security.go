package handler

import (
	"github.com/Dimoonevs/Online_store/internal/models"
	"github.com/Dimoonevs/Online_store/pkg"
	"net/http"
	"strings"
)

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	data := h.service.Security.Test()

	err := pkg.WriteJSON(w, http.StatusOK, data)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) Authentication(w http.ResponseWriter, r *http.Request) {
	admin := &models.Admin{}
	err := pkg.ReadJSON(w, r, admin)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	token, err := h.service.Security.Authentication(admin)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}
	err = pkg.WriteJSON(w, http.StatusOK, token)
	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}
}

func (h *Handler) ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized"})
			return
		}

		token := strings.Split(authorization, "Bearer ")
		if len(token) < 2 {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": "Unauthorized"})
			return
		}

		if err := h.service.Security.ValidateToken(token[1]); err != nil {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
			return
		}
		next.ServeHTTP(w, r)
	})
}

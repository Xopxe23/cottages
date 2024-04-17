package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Xopxe23/cottages/internal/domain"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.WriteErrorResponse(w, err)
		return
	}

	var input domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.WriteErrorResponse(w, err)
		return
	}

	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.WriteErrorResponse(w, err)
		return
	}

	if err = h.auth.CreateUser(r.Context(), input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.WriteErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	h.WriteDataResponse(w, "Registration success")
}

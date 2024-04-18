package rest

import (
	"encoding/json"
	"fmt"
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

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.WriteErrorResponse(w, err)
		return
	}
	var input domain.SignInInput
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
	accessToken, refreshToken, err := h.auth.Authenticate(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.WriteErrorResponse(w, err)
		return
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("refreshToken=%s; HttpOnly", refreshToken))
	h.WriteDataResponse(w, map[string]string{"accessToken": accessToken})
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.WriteErrorResponse(w, err)
		return
	}
	accessToken, refreshToken, err := h.auth.RefreshTokens(r.Context(), cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.WriteErrorResponse(w, err)
		return
	}
	w.Header().Add("Set-Cookie", fmt.Sprintf("refreshToken=%s; HttpOnly", refreshToken))
	h.WriteDataResponse(w, map[string]string{"accessToken": accessToken})
}

package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Xopxe23/cottages/internal/domain"
	"github.com/gorilla/mux"
)

type AuthService interface {
	CreateUser(ctx context.Context, input domain.SignUpInput) error
}

type Handler struct {
	auth AuthService
}

func NewHandler(auth AuthService) *Handler {
	return &Handler{auth}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	}
	return r
}

func (h *Handler) WriteDataResponse(w http.ResponseWriter, data interface{}) {
	resp := domain.ResponseScheme{
		Data: data,
	}
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func (h *Handler) WriteErrorResponse(w http.ResponseWriter, err error) {
	resp := domain.ResponseScheme{
		Error: err.Error(),
	}
	response, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(response)
}

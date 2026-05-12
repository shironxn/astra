package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shironxn/astra/internal/domain"
)

type BookHandler struct {
	service domain.BookService
}

func NewBookHandler(service domain.BookService) domain.BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Book

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := h.service.Create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *BookHandler) GetById(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("id")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := h.service.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("id")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req domain.Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.Update(id, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("id")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

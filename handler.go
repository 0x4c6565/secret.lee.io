package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/0x4c6565/secret.lee.io/pkg/storage"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type textRequestResponse struct {
	Text string `json:"text"`
}

type uuidResponse struct {
	UUID string `json:"uuid"`
}

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	content, err := h.storage.Get(context.Background(), vars["uuid"])
	if err != nil {
		if err == redis.Nil {
			h.handleJSONResponse(w, http.StatusNotFound, nil)
			return
		}

		log.Errorf("Error retrieving secret: %s", err.Error())
		h.handleJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	err = h.storage.Delete(context.Background(), vars["uuid"])
	if err != nil {
		log.Errorf("Error removing secret: %s", err.Error())
		h.handleJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	h.handleJSONResponse(w, http.StatusOK, &textRequestResponse{
		Text: content,
	})
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var req textRequestResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.handleJSONResponse(w, http.StatusBadRequest, nil)
		return
	}
	defer r.Body.Close()

	uuid := uuid.NewString()

	err = h.storage.Set(context.Background(), uuid, req.Text)
	if err != nil {
		log.Errorf("Error storing secret: %s", err.Error())
		h.handleJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	h.handleJSONResponse(w, http.StatusOK, &uuidResponse{
		UUID: uuid,
	})
}

func (h *Handler) handleJSONResponse(w http.ResponseWriter, statusCode int, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(content)
}

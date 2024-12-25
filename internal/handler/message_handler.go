package handler

import (
	"encoding/json"
	"messageOK/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// MessageHandler
type MessageHandler struct {
	useCase usecase.MessageUseCase
}

// Initialize a new message handler
func NewMessageHandler(router *mux.Router, useCase usecase.MessageUseCase) {
	handler := &MessageHandler{useCase: useCase}

	router.HandleFunc("/start", handler.StartSending).Methods("POST")
	router.HandleFunc("/stop", handler.StopSending).Methods("POST")
	router.HandleFunc("/sent", handler.GetSentMessages).Methods("GET")

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

// StartSending : Starting the automatic message sender
// @Summary Start sending messages
// @Description Start the background process to send unsent 2 messages every 2 minutes
// @Tags Messages
// @Success 200 {string} string "OK"
// @Router /start [post]
func (h *MessageHandler) StartSending(w http.ResponseWriter, r *http.Request) {
	go h.useCase.RunAutomaticSender()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Started sending messages."))
}

// StopSending : Stopping the message sender
// @Summary Stop sending messages
// @Description Stop the background process of sending messages
// @Tags Messages
// @Success 200 {string} string "OK"
// @Router /stop [post]
func (h *MessageHandler) StopSending(w http.ResponseWriter, r *http.Request) {
	h.useCase.StopAutomaticSender()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stopped sending messages."))
}

// GetSentMessages : Getting sent messages
// @Summary Get sent messages
// @Description Get a list of messages info that have been sent
// @Tags Messages
// @Success 200 {array} entity.SentMessage
// @Router /sent [get]
func (h *MessageHandler) GetSentMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := h.useCase.GetSentMessages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

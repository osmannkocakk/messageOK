package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"messageOK/config"
	"messageOK/internal/entity"
	"messageOK/internal/repository"
	"messageOK/pkg/redis"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type MessageUseCase interface {
	RunAutomaticSender()
	StopAutomaticSender()
	GetSentMessages() ([]entity.SentMessage, error)
}

type messageUseCase struct {
	repo        repository.MessageRepository
	redisClient redis.RedisClient
	config      *config.Config
	stopChan    chan struct{}
}

func NewMessageUseCase(repo repository.MessageRepository, rc redis.RedisClient, cfg *config.Config) MessageUseCase {
	return &messageUseCase{
		repo:        repo,
		redisClient: rc,
		config:      cfg,
		stopChan:    make(chan struct{}),
	}
}

func (uc *messageUseCase) RunAutomaticSender() {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			messages, err := uc.repo.GetUnsentMessages()
			if err != nil {
				log.Println("Error while fetching messages:", err)
				continue
			}

			for _, msg := range messages {
				if len(msg.Content) > 160 {
					msg.Content = msg.Content[:160] // Truncate to 160 characters
				}

				err = uc.sendMessage(msg)
				if err != nil {
					log.Println("Error sending message:", err)
					continue
				}

				err = uc.repo.MarkMessageAsSent(msg.ID)
				if err != nil {
					log.Println("Error making message as sent:", err)
					continue
				}
			}
		case <-uc.stopChan:
			log.Println("Stopping automatic sender...")
			return
		}
	}
}

func (uc *messageUseCase) StopAutomaticSender() {
	close(uc.stopChan)
}

func (uc *messageUseCase) sendMessage(msg entity.Message) error {
	payload := map[string]string{
		"to":      msg.To,
		"content": msg.Content,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", uc.config.Message.ApiUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(uc.config.Message.HeaderKey, uc.config.Message.HeaderValue)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send message")
	}

	messageID := uuid.New().String()

	responseData := map[string]string{
		"message":   "Accepted",
		"messageId": messageID,
	}

	var response struct {
		Message   string `json:"message"`
		MessageID string `json:"messageId"`
	}

	response.MessageID = responseData["messageId"]

	// Redis Cache for sent message
	uc.redisClient.Set(response.MessageID, time.Now().UTC().String())

	return nil
}

func (uc *messageUseCase) GetSentMessages() ([]entity.SentMessage, error) {
	return uc.redisClient.GetSentMessages()
}

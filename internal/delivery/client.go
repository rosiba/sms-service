package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sms-service/internal/model"
	"time"
)

const DeliveryAccepted = "Accepted"

type DeliveryRequest struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type DeliveryResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

func deliver(message model.Message, resultChan chan DeliveryResult, errChan chan error) {
	c := &http.Client{
		Timeout: 2 * time.Second,
	}

	url := os.Getenv("DELIVERY_URL")
	if url == "" {
		errChan <- errors.New("DELIVERY_URL environment variable not set")
		return
	}

	reqBody, err := json.Marshal(DeliveryRequest{
		To:      message.Recipient,
		Content: message.Content,
	})
	if err != nil {
		errChan <- fmt.Errorf("failed to unmarshal request body: %v", err)
		return
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		errChan <- fmt.Errorf("failed to create request: %v", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-ins-auth-key", os.Getenv("DELIVERY_AUTH_KEY"))

	resp, err := c.Do(request)
	if err != nil {
		errChan <- fmt.Errorf("failed to send request: %v", err)
		return
	}
	if resp.StatusCode != http.StatusAccepted {
		errChan <- fmt.Errorf("failed to deliver message: status %s", resp.Status)
		return
	}

	respBody := DeliveryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		errChan <- fmt.Errorf("failed to decode response body: %v", err)
		return
	}
	if respBody.Message != DeliveryAccepted {
		errChan <- fmt.Errorf("failed to deliver message: status %s", respBody.Message)
		return
	}

	resultChan <- DeliveryResult{
		MessageID:  message.ID,
		ExternalID: respBody.MessageID,
	}
}

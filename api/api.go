package api

const (
	DeliveryStartSuccessful = "Message delivery started successfully"
	DeliveryStopSuccessful  = "Message delivery stopped successfully"
	InvalidAction           = "Invalid action"
	InvalidRequestBody      = "Invalid request body"
	FailedSaveMessage       = "Failed to save message"
	MessageCreated          = "Message created with pending status"
)

type Response struct {
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Error        error  `json:"error,omitempty"`
}

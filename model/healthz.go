package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	message string `json:"message"`
}

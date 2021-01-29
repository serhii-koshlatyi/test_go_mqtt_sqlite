package model

// MQTTMessage model that we get in mqtt message
type MQTTMessage struct {
	ID      int    `json:"_" db:"id"`
	Message string `json:"message" db:"message"`
	Name    string `json:"name" db:"message"`
}

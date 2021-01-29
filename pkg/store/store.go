package store

import (
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/model"
)

// Store represent interface for database store
type Store interface {
	Create(model model.MQTTMessage) error
}

package sqlitestore

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Driver for sqlite

	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/config"
	"github.com/sergeykoshlatuu/test_go_mqtt_sqlite/pkg/model"
)

const (
	MessageTable = "messages"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewStore(cfg config.SQLite) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", cfg.DSN)
	if err != nil {
		return nil, err
	}

	store := &SQLiteStore{
		db: db,
	}

	return store, nil
}

func (s *SQLiteStore) Create(model model.MQTTMessage) error {
	_, err := s.db.Exec("insert into messages (name, message) values  ($1, $2)", model.Name, model.Message)
	if err != nil {
		return err
	}

	return nil
}

package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Connection string `json:"connection"`
	Driver     string `json:"driver"`
}

type Storage struct {
	db *sql.DB
}

func New(cfg *Config) (*Storage, error) {
	connection, err := sql.Open(cfg.Driver, cfg.Connection)

	if err != nil {
		return nil, err
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	stmt, err := connection.Prepare(SQLStringCreateTableJobless)

	if err != nil {
		return nil, err
	}

	stmt.Exec()

	stmt, err = connection.Prepare(SQLStringCreateTableJobgivers)

	if err != nil {
		return nil, err
	}

	stmt.Exec()

	stmt, err = connection.Prepare(SQLStringCreateTableArchive)

	if err != nil {
		return nil, err
	}

	stmt.Exec()

	return &Storage{db: connection}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

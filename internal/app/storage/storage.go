package storage

import (
	"database/sql"
	"webserver/internal/app/models"

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

func (s *Storage) ExecuteString(cmd string, a ...interface{}) error {

	q, err := s.db.Prepare(cmd)

	if err != nil {
		return err
	}

	defer q.Close()

	if _, err := q.Exec(a...); err != nil {
		return err
	}

	return nil
}

func (s *Storage) InsertJobless(jobless *models.Jobless) error {

	return s.ExecuteString(SQLStringInsertJobless,
		jobless.LastName,
		jobless.FirstName,
		jobless.Patronymic,
		jobless.Age,
		jobless.Passport,
		jobless.PassportDate,
		jobless.Region,
		jobless.Address,
		jobless.Phone,
		jobless.StudyPlace,
		jobless.StudyAddress,
		jobless.StudyType,
		jobless.Registrar,
		jobless.RegDate,
		jobless.Payment,
		jobless.Comment,
	)
}

func (s *Storage) InsertInArchive(archive *models.Archive) error {
	return s.ExecuteString(SQLStringInsertArchive,
		archive.JoblessId,
		archive.JobId,
		archive.Date,
		archive.Archivist,
	)
}

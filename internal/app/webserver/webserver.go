package webserver

import (
	"webserver/internal/app/configurator"
	"webserver/internal/app/storage"
)

type WEBServer struct {
	Config  *configurator.Config
	Storage *storage.Storage
}

func New() *WEBServer {
	return &WEBServer{}
}

func (s *WEBServer) Start() error {
	var err error

	if s.Config, err = configurator.Open("./config.json"); err != nil {
		return err
	}

	if s.Storage, err = storage.New(&s.Config.DBConfig); err != nil {
		return err
	}

	defer s.Storage.Close()

	return nil
}

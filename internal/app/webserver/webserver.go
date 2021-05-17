package webserver

import (
	"net/http"
	"webserver/internal/app/configurator"
	"webserver/internal/app/storage"

	"github.com/gorilla/mux"
)

type Config struct {
	Address string `json:"address"`
}

type WEBServer struct {
	Config  *configurator.Config
	Storage *storage.Storage
	Router  *mux.Router
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

	s.Router = mux.NewRouter()

	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	if err := http.ListenAndServe(s.Config.WSAddress, s.Router); err != nil {
		return err
	}

	return nil
}

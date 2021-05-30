package webserver

import (
	"html/template"
	"log"
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

type ButtonTemplate struct {
	Text string
	URL  string
}

type MenuTemplate struct {
	Header  string
	Buttons []ButtonTemplate
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

	s.Router.HandleFunc("/", s.RedirectToIndex)
	s.Router.HandleFunc("/index", s.IndexHandler)

	s.Router.HandleFunc("/vacations", s.MenuHandler(MenuTemplate{
		Header: "Работа с вакансиями",
		Buttons: []ButtonTemplate{
			ButtonTemplate{Text: "Добавить", URL: "/add/vacations"},
			ButtonTemplate{Text: "Изменить", URL: "/change/vacations"},
			ButtonTemplate{Text: "Удалить", URL: "/remove/vacations"},
			ButtonTemplate{Text: "Просмотреть", URL: "/show/vacations"},
			ButtonTemplate{Text: "В меню", URL: "/index"},
		},
	}))

	s.Router.HandleFunc("/jobless", s.MenuHandler(MenuTemplate{
		Header: "Работа с безработными",
		Buttons: []ButtonTemplate{
			ButtonTemplate{Text: "Добавить", URL: "/add/jobless"},
			ButtonTemplate{Text: "Изменить", URL: "/change/jobless"},
			ButtonTemplate{Text: "Удалить", URL: "/remove/jobless"},
			ButtonTemplate{Text: "Просмотреть", URL: "/show/jobless"},
			ButtonTemplate{Text: "В меню", URL: "/index"},
		},
	}))

	s.Router.HandleFunc("/archive", s.MenuHandler(MenuTemplate{
		Header: "Работа с архивом",
		Buttons: []ButtonTemplate{
			ButtonTemplate{Text: "Добавить", URL: "/add/archive"},
			ButtonTemplate{Text: "Просмотреть", URL: "/show/archive"},
			ButtonTemplate{Text: "В меню", URL: "/index"},
		},
	}))

	if err := http.ListenAndServe(s.Config.WSAddress, s.Router); err != nil {
		return err
	}

	return nil
}

func (s *WEBServer) RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func (s *WEBServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/templates/index.html")

	if err != nil {
		log.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func (s *WEBServer) MenuHandler(content MenuTemplate) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./html/templates/menu.html")

		if err != nil {
			log.Println(err)
		} else {
			t.Execute(w, &content)
		}
	}
}

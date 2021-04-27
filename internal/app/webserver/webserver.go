package webserver

type WEBServer struct{}

func New() *WEBServer {
	return &WEBServer{}
}

func (s *WEBServer) Start() error {
	return nil
}

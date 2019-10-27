package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type WsServer struct {
	http.Server
	websocket.Upgrader
}

func NewWsServer(addr string) *WsServer {
	srv := &WsServer{
		Server: http.Server{
			Addr: addr,
		},
		Upgrader: websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
	}
	srv.setupRoutes()
	return srv
}

func (s *WsServer) setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/events/ws", s.eventHandler)
	s.Handler = r
}

func (s *WsServer) eventHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Str("remote", r.RemoteAddr).Msg("Serving ws request")
	conn, err := s.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Err(err).Str("remote", r.RemoteAddr).Msg("Bad Request")
		return
	}
	go s.serve(conn)
}

func (s *WsServer) serve(conn *websocket.Conn) {
	for {
		msg := &Msg{}
		if err := conn.ReadJSON(msg); err != nil {
			log.Error().Err(err).Msg("Json read error")
			conn.Close()
		}
	}
}

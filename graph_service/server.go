package main

import (
	"net/http"
	"github.com/dgraph-io/dgo"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
)

type Server struct {
	http.Server
	dgraph *dgo.Dgraph	
}

type handlerFunc func(ctx *requestCtx) error

type requestCtx struct {
	logger *zerolog.Logger
	r *http.Request

	code int
	resp []byte
}

func (ctx *requestCtx) writeResp(w http.ResponseWriter) error {
	w.WriteHeader(ctx.code)
	_, err := w.Write(ctx.resp)
	return err
}

func (s *Server) setupHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/db/graph/student", middleware(s.graphHandler))
	r.HandleFunc("/db/course/students", middleware(s.courseStudentsHandler))
	r.HandleFunc("/db/student/info", middleware(s.studentInfoHandler))
	r.HandleFunc("/db/tests/questions", middleware(s.testQuestionsHandler))
	s.Server.Handler = r
}

func middleware(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With().
			Str("path", r.RequestURI).
			Str("remote", r.RemoteAddr).
			Logger()
		logger.Info().Msg("Serving request")
		ctx := &requestCtx{
			logger: &logger,
			r: r,
			code: 200,
		}
		if err := h(ctx); err != nil {
			log.Error().Err(err).Msg("Request failed")
		} else {
			log.Info().Msg("Request succeeded")
		}
		if err := ctx.writeResp(w); err != nil {
			log.Error().Err(err).Msg("Failed to write response")
		}
	}
}

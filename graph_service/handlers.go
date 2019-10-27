package main

import (
	"context"
	"errors"
	"encoding/json"
)

var (
	errBadRequest = errors.New("bad request")
)

func (s *Server) graphHandler(ctx *requestCtx) error {
	user := ctx.r.FormValue("name")
	if user == "" {
		ctx.code = 400
		return errBadRequest
	}
	ctx.logger.Info().Str("user", user).Msg("Get graph for user")
	query := `
		query select_by_name($name: string) {
			user_graph(func: eq(name, $name)) {
				studied_course {
					name
					has_topic {
						name
					}
				}
				passed_topic @facets(strength) {
					name
				}
			}	
		}
		`
	txn := s.dgraph.NewTxn()
	defer txn.Discard(context.TODO())
	res, err := txn.QueryWithVars(context.TODO(), query, map[string]string{"$name": user})
	if err != nil {
		ctx.code = 500
		return err
	}
	ctx.logger.Info().
		Str("user", user).
		RawJSON("graph", res.Json).
		Msg("Got graph for user")
	graph := UserGraph{}
	if err = json.Unmarshal(res.Json, &graph); err != nil {
		ctx.code = 500
		return err
	}
	list := graph.transformToList()
	marshaled, err := json.Marshal(list)
	if err != nil {
		ctx.code = 500
		return err
	}
	ctx.resp = marshaled
	return nil
}

func (s *Server) courseStudentsHandler(ctx *requestCtx) error {
	course := ctx.r.FormValue("course")
	if course == "" {
		ctx.code = 400
		return errBadRequest
	}
	ctx.logger.Info().Str("course", course).Msg("Select students for course")
	query := `
		query select_by_name($name: string) {
			student_list(func: eq(name, $name)) {
				~studied_course {
					name
					rating
					warning
				}
  			}
		}
	`
	txn := s.dgraph.NewTxn()
	defer txn.Discard(context.TODO())
	res, err := txn.QueryWithVars(context.TODO(), query, map[string]string{"$name": course})
	if err != nil {
		ctx.code = 500
		return err
	}
	ctx.logger.Info().
		Str("course", course).
		RawJSON("students", res.Json)
	ctx.resp = res.Json
	return nil
}

func (s *Server) studentInfoHandler(ctx *requestCtx) error {
	student := ctx.r.FormValue("student")
	if student == "" {
		ctx.code = 400
		return errBadRequest 
	}
	ctx.logger.Info().Str("student", student).Msg("Select data for student")
	query := `
		query select_by_name($name: string) {
			student_info(func: eq(name, $name)) {
				name
				rating
				warning
				courses: count(studied_course)
				topics: count(passed_topic)
  			}
		}
	`
	txn := s.dgraph.NewTxn()
	defer txn.Discard(context.TODO())
	res, err := txn.QueryWithVars(context.TODO(), query, map[string]string{"$name": student})
	if err != nil {
			ctx.code = 500
			return err
	}
	ctx.logger.Info().
		Str("student", student).
		RawJSON("info", res.Json).
		Msg("got result for user")
	ctx.resp = res.Json
	return nil
}

func (s *Server) testQuestionsHandler(ctx *requestCtx) error {
	return nil
}

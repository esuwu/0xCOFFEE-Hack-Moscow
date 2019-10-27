package main

import (
	"context"
	"flag"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	dgraph := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	demoInit := flag.Bool("init", false, "init db")
	flag.Parse()
	if *demoInit {
		initDemoDB(dgraph)
	}
	srv := Server{
		Server: http.Server{Addr: ":8888"},
		dgraph: dgraph,
	}
	srv.setupHandlers()
	srv.ListenAndServe()
}

func initDemoDB(dgraph *dgo.Dgraph) {
	op := &api.Operation{
		Schema: `type User {
  					name: string
					rating: int
					warning: bool
  					type: string
					studied_course: [Course]
					passed_topic: [Topic]
  					has_result: [TestResult]
  					has_activity: [Activity]
				}

				type Course {
  					name: string
					teacher: string
  					has_topic: [Topic]
  					related_to: [Course]
				}

				type Topic {
  					name: string
  					from_course: [Course]
  					related_to: [Topic]
				}

				type Test {
					question: string
					answer: string
					for_topic: [Topic]
				}

				type TestResult {
					scores: string
  					for_test: [Test]
					for_user: [User]
				}

				type Activity {
  					of_user: [User]
  					on_topic: [Topic]
				}

				name: string @index(exact, term) .

				studied_course: [uid] @reverse .
				passed_topic: [uid] .
				has_results: [uid] .
				has_activity: [uid] .
				has_topic: [uid] .
				related_to: [uid] .
				for_topic: [uid] .
				on_topic: [uid] .
				`,
	}
	err := dgraph.Alter(context.TODO(), op)
	if err != nil {
		panic(err)
	}
	mu := api.Mutation{}
	mu.SetNquads = []byte(`
	_:course1 <name> "Math" .
    _:course1 <dgraph.type> "Course" .
    _:course2 <name> "Physics" .
    _:course2 <dgraph.type> "Course" .
    _:course2 <related_to> _:course1 .
    
    _:topic1 <name> "Linear Algebra" .
    _:topic1 <dgraph.type> "Topic" .
    _:topic2 <name> "Real Analysis" .
    _:topic2 <dgraph.type> "Topic" .
    
    _:topic3 <name> "Mechanics" .
    _:topic3 <dgraph.type> "Topic" .
		_:topic3 <related_to> _:topic2 .
    _:topic4 <name> "Electrodynamics" .
    _:topic4 <dgraph.type> "Topic" .
    _:topic4 <related_to> _:topic2 .

    _:course1 <has_topic> _:topic1 .
    _:course1 <has_topic> _:topic2 .
    _:course2 <has_topic> _:topic3 .
    _:course2 <has_topic> _:topic4 .

    
    _:user1 <name> "Alex" .
    _:user1 <dgraph.type> "User" .
    _:user1 <type> "student" .
    _:user1 <rating> "70" .
    _:user1 <warning> "false" .
    _:user1 <studied_course> _:course1 .
    _:user1 <studied_course> _:course2 .
    _:user1 <passed_topic> _:topic1 (strength=10) .
    _:user1 <passed_topic> _:topic4 (strength=5) .`)
	txn := dgraph.NewTxn()
	_, err = txn.Mutate(context.TODO(), &mu)
	if err != nil {
		panic(err)
	}
	if err := txn.Commit(context.TODO()); err != nil {
		panic(err)
	}
}

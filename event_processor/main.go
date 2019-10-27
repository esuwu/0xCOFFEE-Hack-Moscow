package main

func main() {
	srv := NewWsServer(":8889")
	panic(srv.ListenAndServe())
}

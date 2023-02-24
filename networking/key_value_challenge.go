package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type DB struct {
	m sync.Map
}

func (db *DB) Get(key string) []byte {
	val, ok := db.m.Load(key)
	if !ok {
		return nil
	}

	return val.([]byte)
}

func (db *DB) Set(key string, value []byte) {
	db.m.Store(key, value)
}

type Server struct {
	db DB
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postHandler(w, r)
		return
	case http.MethodGet:
		s.getHandler(w, r)
		return
	}
	http.Error(w, "bad method", http.StatusMethodNotAllowed)
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // trim leading /
	data := s.db.Get(key)
	if data == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Write(data)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:] // trim leading /
	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, 1<<10)
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}
	s.db.Set(key, data)
	fmt.Fprintf(w, "%s set\n", key)
}

func main() {
	var s Server
	http.HandleFunc("/", s.Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
curl -d 'hello' http://localhost:8080/k1
curl http://localhost:8080/k1
curl -i http://localhost:8080/k2

Limit value size to 1k
*/

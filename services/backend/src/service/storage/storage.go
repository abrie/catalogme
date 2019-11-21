package storage

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type Storage struct {
	Directory string
}

func (s *Storage) Datafile() string {
	return path.Join(s.Directory, "schema.json")
}

func (s *Storage) GetHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(s.Datafile())
	if err != nil {
		log.Printf("Failed to read the file %s: %s", s.Datafile(), err.Error())
		http.Error(w, "Failed to load.", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

package api

import (
	"encoding/json"
	"net/http"
)

// optional code omitted

type Server struct{}

// GetPetById implements ServerInterface.
func (s Server) GetPetById(w http.ResponseWriter, r *http.Request, petId int) {
	panic("unimplemented")
}

func NewServer() Server {
	return Server{}
}

// (GET /ping)
func (Server) GetPing(w http.ResponseWriter, r *http.Request) {
	resp := Pong{
		Ping: "pong",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

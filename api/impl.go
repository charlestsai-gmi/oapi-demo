package api

import (
	"encoding/json"
	"net/http"

	"github.com/charlestsai-gmi/oapi-demo/common"
)

// optional code omitted

type Server struct{}

// GetPetById implements ServerInterface.
func (s Server) GetPetById(w http.ResponseWriter, r *http.Request, petId int) {
	resp := common.Pet{
		Id:   petId,
		Name: "dog",
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
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

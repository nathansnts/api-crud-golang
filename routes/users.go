package routes

import (
	"estudos/golang/handlers"
	"net/http"
)

//Rotas de acesso

func RegistrarRotas(mux *http.ServeMux) {
	// Rotas para os usuários
	mux.HandleFunc("/users", handlers.HandlerUsers)
	mux.HandleFunc("/users/", handlers.HandlerUserID)
}

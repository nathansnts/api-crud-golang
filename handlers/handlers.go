package handlers

import (
	"encoding/json"
	"estudos/golang/model"
	"estudos/golang/storage"
	"net/http"
	"strconv"
)

// Define o método durante requisição na api
func HandlerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//Retorna todos os usuários disponíveis
		listUsers(w)

	case http.MethodPost:
		//Add novo usuário
		createUser(w, r)
	default:
		http.Error(w, "Método não suportado para essa requisição", http.StatusMethodNotAllowed)
	}
}

// Obter usuários pelo ID
func HandlerUserID(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idstr)
	if err != nil || id < 1 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return

	}

	user, found := storage.GetUser(id)

	if !found {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Listar usuários
func listUsers(w http.ResponseWriter) {
	users := storage.ListUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Criar usuários
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"mensagem": "Usuário criado com sucesso",
	}

	storage.CreateUser(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

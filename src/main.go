package main

import (
	"estudos/golang/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Registro de rotas
	routes.RegistrarRotas(mux)

	fmt.Println("Server ins running in port 8080...")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

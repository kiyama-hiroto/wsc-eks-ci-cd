package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Defina uma struct para representar a estrutura do seu JSON
type User struct {
	Name      string `json:"name"`
	Ocupacao string `json:"ocupacao"`
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "OK")
}

func v1(w http.ResponseWriter, req *http.Request) {
	// Crie uma instância da sua struct com os dados desejados
	user := User{
		Name:      "Gemini",
		Ocupacao: "Desenvolvedor de IA",
	}

	// Defina o cabeçalho Content-Type para indicar que a resposta é um JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifique a struct em JSON e envie para o ResponseWriter
	// json.NewEncoder() cria um encoder que escreve diretamente no ResponseWriter
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/v1", v1)
	http.ListenAndServe(":8000", nil)
}
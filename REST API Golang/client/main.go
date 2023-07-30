package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func main() {
	//Verificação do endpoint
	response, error := http.Get("http://localhost:8080/users")
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	//Se o status code é o esperado
	if response.StatusCode != 200 {
		fmt.Println("Not Success", response.StatusCode)
		return
	}

	//Lê o body da requisição
	body, error := io.ReadAll(response.Body)

	//Popula a variável com o retorno esperado
	var resp []User
	error = json.Unmarshal(body, &resp)

	if error != nil {
		fmt.Println("Erro when getting data", error.Error())
		return
	}

	fmt.Println(resp)
}

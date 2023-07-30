package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	Id   int
	Name string
}

func getUsers(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	writer.Header().Set("Content=Type", "application/json")
	json.NewEncoder(writer).Encode([]User{
		{
			Id:   1,
			Name: "Matheus",
		},
		{
			Id:   2,
			Name: "Billy",
		},
	})
}

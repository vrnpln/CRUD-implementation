package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string //обновляем переменную task через POST-запрос, а затем возвращаем обновленное значение при GET-запросе

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", task)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	json.NewDecoder(r.Body).Decode(&reqBody)

	task = reqBody.Message
	fmt.Fprintln(w, task)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}

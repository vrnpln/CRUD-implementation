package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message Message

	// Декодирование JSON в структуру Message
	json.NewDecoder(r.Body).Decode(&message)
	// Сохранение в БД
	DB.Create(&message)
	// Кодирирование структуры Message в JSON и отправка
	json.NewEncoder(w).Encode(message)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var messages []Message

	// Получение всех записей из БД
	DB.Find(&messages)
	// Кодирование в JSON и отправка
	json.NewEncoder(w).Encode(messages)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}

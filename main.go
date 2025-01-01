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

func UpdateMessages(w http.ResponseWriter, r *http.Request) {
	var message Message

	// Получение переменных из маршрута
	vars:= mux.Vars(r)
	// Получение ID из переменных маршрута
	id := vars["id"]
	// Поиск задачи по ID
	DB.First(&message,id)

	// Кодирирование структуры Message в JSON и отправка
	json.NewDecoder(r.Body).Decode(&message)
	// Сохранение изменений в БД
	DB.Save(&message)
	// Кодирование в JSON и отправка обновленной задачи
	json.NewEncoder(w).Encode(message)
}

func DeleteMessages( w http.ResponseWriter, r *http.Request) {
	var message Message

	// Получение переменных из маршрута
	vars:= mux.Vars(r)
	// Получение ID из переменных маршрута
	id := vars["id"]
	// Поиск задачи по ID
	DB.First(&message,id)

	// Удаление найденной задачи в БД
	DB.Delete(&message)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/{id}", UpdateMessages).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessages).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

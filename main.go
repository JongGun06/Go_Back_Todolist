package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todolist_back_one/handlers"
	"todolist_back_one/routes"
)

// MongoDB конфигурация
const mongoURI string =  "mongodb+srv://nagidevfullstack:Hadi2017g@golangdb.ebvlakf.mongodb.net/?retryWrites=true&w=majority&appName=GolangDB"

// Middleware для CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		// Устанавливаем заголовки CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		// Логируем запрос для отладки
		log.Printf("Получен запрос: %s %s", r.Method, r.URL.Path)

		// Обрабатываем предварительные запросы OPTIONS
		if r.Method == "OPTIONS" {
			log.Println("Обработан запрос OPTIONS")
			w.WriteHeader(http.StatusOK)
			return
		}

		// Передаём запрос следующему обработчику
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Подключение к MongoDB Atlas
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Проверка подключения
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB Atlas!")

	// Установка клиента MongoDB в handlers
	handlers.SetClient(client)

	// Создание маршрутизатора
	router := mux.NewRouter()

	// Настройка маршрутов
	routes.SetupRoutes(router)

	// Добавляем middleware для CORS
	corsRouter := enableCORS(router)

	// Запуск сервера
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
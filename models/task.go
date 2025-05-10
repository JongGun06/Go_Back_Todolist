package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task представляет структуру задачи
type Task struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
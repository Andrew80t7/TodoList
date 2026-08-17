package http

import (
	"encoding/json"
	"errors"
	"time"
)

//Data Transfer Object - структуры, чтобы передавать какие-то данные

type CompletedTaskDTO struct {
	Completed bool `json:"completed"`
}

// Нужна, чтобы принять входящий http запрос
type TaskDTO struct {
	Title       string
	Description string
}

func (t *TaskDTO) ValidateToCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {

	// Ошибка в структуре может быть, например канал нельзя в байты,
	//также рекурсивные вызовы функции внутри структуры
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)

}

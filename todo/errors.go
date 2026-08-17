package todo

import "errors"

// Задача не найдена
var ErrTaskNotFound = errors.New("task not found")

// Задача уже создана
var ErrTaskAlreadyExists = errors.New("task already exists")

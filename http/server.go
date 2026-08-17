package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	httpHandlers *Handlers
}

func NewServer(httpHandlers *Handlers) *Server {
	return &Server{
		httpHandlers: httpHandlers,
	}
}

// Routing
func (s *Server) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandlerCreateTask)

	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetTask)

	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetAllTasks)

	router.Path("/tasks").Methods("GET").
		Queries("completed", "true").
		HandlerFunc(s.httpHandlers.HandlerGetAllUncompletedTasks)

	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlerCompleteTask)

	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandlerDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}

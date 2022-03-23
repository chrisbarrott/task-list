/*

// THIS CREATES TASKS IN A WEB SERVER, BUT COULDNT DO MUCH MORE THAN RUN THEM IN A JSON FILE

*/

package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// task object
type Task struct {
	ID     uuid.UUID `json:"id"`
	Task   string    `json:"task"`
	Status string    `json:"status"`
}

type Server struct {
	*mux.Router
	tasks []Task
}

// create a router for localhost
func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		tasks:  []Task{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/tasks", s.listTasks()).Methods("GET")
	s.HandleFunc("/tasks", s.createTask()).Methods("POST")
	s.HandleFunc("/tasks/{id}", s.completeTask()).Methods("DELETE")
	//s.HandleFunc("/tasks/{id}", s.updateTask()).Methods("PUT") // not worked this out yet
}

func (s *Server) createTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t Task

		// handle a bad request
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// creates a new UUID
		t.ID = uuid.New()
		s.tasks = append(s.tasks, t)

		// set the header
		w.Header().Set("Content-Type", "application/json")

		// validate the response
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set the header
		w.Header().Set("Content-Type", "application/json")

		// handle broken server
		if err := json.NewEncoder(w).Encode(s.tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Server) completeTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]

		// parse the ID
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// loop round each known task and returnthem all
		for i, task := range s.tasks {
			if task.ID == id {
				s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
				break
			}
		}
		// no response, just give 200 response
	}
}

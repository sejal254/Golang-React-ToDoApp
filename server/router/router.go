package router

import (
	//"golang-react-todo/middleware"

	"github.com/gorilla/mux"
	"github.com/sejal254/Go-React-ToDoApp/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deteleAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return router

}

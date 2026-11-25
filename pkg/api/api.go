package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const dateFmt = "20060102"

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/nextdate", nextDateHandler)
	mux.HandleFunc("GET /api/task", getTaskHandler)
	mux.HandleFunc("POST /api/task", addTaskHandler)
	mux.HandleFunc("PUT /api/task", putTaskHandler)
	mux.HandleFunc("DELETE /api/task", delTaskHandler)
	mux.HandleFunc("POST /api/task/done", doneHandler)
	mux.HandleFunc("GET /api/tasks", tasksHandler)
	return mux
}

func writeJson(w http.ResponseWriter, code int, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if _, err := w.Write(resp); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

func errorJson(w http.ResponseWriter, code int, msg string) {
	writeJson(w, code, map[string]string{"error": msg})
}

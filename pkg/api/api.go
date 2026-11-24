package api

import (
	"encoding/json"
	"net/http"
)

const dateFmt = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/task/done", doneHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
}

func writeJson(w http.ResponseWriter, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(resp)
}

func errorJson(w http.ResponseWriter, msg string) {
	writeJson(w, map[string]string{"error": msg})
}

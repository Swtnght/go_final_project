package api

import (
	"net/http"

	"github.com/Swtnght/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		errorJson(w, err.Error())
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

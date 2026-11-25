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
		errorJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJson(w, http.StatusOK, TasksResp{
		Tasks: tasks,
	})
}

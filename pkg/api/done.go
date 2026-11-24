package api

import (
	"net/http"
	"time"

	"github.com/Swtnght/go_final_project/pkg/db"
)

func doneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		errorJson(w, "Не указан идентификатор")
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		errorJson(w, "Задача не найдена")
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			errorJson(w, err.Error())
			return
		}
		writeJson(w, map[string]any{})
		return
	}

	date, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		errorJson(w, err.Error())
		return
	}

	err = db.UpdateDate(date, id)
	if err != nil {
		errorJson(w, err.Error())
		return
	}
	writeJson(w, map[string]any{})
}

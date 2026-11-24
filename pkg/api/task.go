package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Swtnght/go_final_project/pkg/db"
)

type taskReq struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type taskResp struct {
	ID string `json:"id"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		putTaskHandler(w, r)
	case http.MethodDelete:
		delTaskHandler(w, r)
	default:
		errorJson(w, "неподдерживаемый метод запроса")
	}
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	writeJson(w, task)
}

func validateTask(req *taskReq) error {
	if req.Title == "" {
		return errors.New("задача должна содержать заголовок")
	}

	now := time.Now()
	today := now.Format(dateFmt)
	if req.Date == "" {
		req.Date = today
	} else {
		_, err := time.Parse(dateFmt, req.Date)
		if err != nil {
			return errors.New("некорректная дата")
		}
	}

	if req.Repeat != "" {
		next, err := NextDate(now, req.Date, req.Repeat)
		if err != nil {
			return err
		}
		if req.Date < today {
			req.Date = next
		}
	} else if req.Date < today {
		req.Date = today
	}
	return nil
}

func delTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		errorJson(w, "Не указан идентификатор")
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		errorJson(w, "Задача не найдена")
		return
	}
	writeJson(w, map[string]any{})
}

func putTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req taskReq
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r.Body); err != nil {
		errorJson(w, "некорректный запрос")
		return
	}
	err := json.Unmarshal(buf.Bytes(), &req)
	if err != nil {
		errorJson(w, "некорректный запрос")
		return
	}

	if req.ID == "" {
		errorJson(w, "Не указан идентификатор")
		return
	}

	err = validateTask(&req)
	if err != nil {
		errorJson(w, err.Error())
		return
	}

	task := db.Task{
		ID:      req.ID,
		Date:    req.Date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}
	err = db.UpdateTask(&task)
	if err != nil {
		errorJson(w, "Задача не найдена")
		return
	}
	writeJson(w, map[string]any{})
}

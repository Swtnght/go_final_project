package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Swtnght/go_final_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	err = validateTask(&req)
	if err != nil {
		errorJson(w, err.Error())
		return
	}

	id, err := db.Insert(req.Date, req.Title, req.Comment, req.Repeat)
	if err != nil {
		errorJson(w, err.Error())
		return
	}
	writeJson(w, taskResp{ID: strconv.FormatInt(id, 10)})
}

package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const maxInterval = 400

func afterNow(date, now time.Time) bool {
	day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return date.After(day)
}

func NextDate(now time.Time, dstart, repeat string) (string, error) {
	date, err := time.Parse(dateFmt, dstart)
	if err != nil {
		return "", errors.New("некорректная дата")
	}

	parts := strings.Split(repeat, " ")
	switch parts[0] {
	case "y":
		if len(parts) != 1 {
			return "", errors.New("некорректный формат правила")
		}
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	case "d":
		if len(parts) != 2 {
			return "", errors.New("не указан интервал в днях")
		}
		interval, err := strconv.Atoi(parts[1])
		if err != nil || interval < 1 || interval > maxInterval {
			return "", errors.New("некорректный интервал в днях")
		}
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
	case "w", "m":
		return "", errors.New("формат повторения не поддерживается")
	default:
		return "", errors.New("неизвестный формат правила")
	}
	return date.Format(dateFmt), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	var now time.Time
	var err error
	nowStr := r.FormValue("now")
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFmt, nowStr)
		if err != nil {
			http.Error(w, "некорректный параметр now", http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := fmt.Fprint(w, next); err != nil {
		log.Printf("ошибка записи ответа: %v", err)
	}
}

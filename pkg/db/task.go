package db

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func Tasks(limit int) ([]*Task, error) {
	rows, err := db.Query(`SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit`,
		sql.Named("limit", limit),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}
	for rows.Next() {
		t := &Task{}
		var id int64
		err := rows.Scan(&id, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		t.ID = strconv.FormatInt(id, 10)
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	t := &Task{}
	var tid int64
	err := db.QueryRow(`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id`,
		sql.Named("id", id),
	).Scan(&tid, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return nil, err
	}
	t.ID = strconv.FormatInt(tid, 10)
	return t, nil
}

func Insert(date, title, comment, repeat string) (int64, error) {
	res, err := db.Exec(
		`INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`,
		sql.Named("date", date),
		sql.Named("title", title),
		sql.Named("comment", comment),
		sql.Named("repeat", repeat),
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateTask(task *Task) error {
	res, err := db.Exec(
		`UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`,
		sql.Named("id", task.ID),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for updating task")
	}
	return nil
}

func UpdateDate(date, id string) error {
	res, err := db.Exec(`UPDATE scheduler SET date = :date WHERE id = :id`,
		sql.Named("date", date),
		sql.Named("id", id),
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for updating task")
	}
	return nil
}

func DeleteTask(id string) error {
	res, err := db.Exec(`DELETE FROM scheduler WHERE id = :id`,
		sql.Named("id", id),
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for deleting task")
	}
	return nil
}

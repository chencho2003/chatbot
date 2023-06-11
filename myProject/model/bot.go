package model

import (
	"myProject/datastore/postgres"
	"database/sql"
	"github.com/pkg/errors"
	"fmt"
)

type Bot struct {
	Id int `json:"id"`
	Question string `json:"question"`
	Answer string `json:"answer"`
	ProfileImage string `json:"profile_image"`
}
const (
	queryQNA   = "INSERT INTO bot(question,answer) VALUES($1, $2);"
	deleteQNA = "DELETE FROM bot WHERE id = $1;"
	accessingQNA = "SELECT answer FROM bot WHERE question ILIKE $1 LIMIT 1"
	updateQNA = "UPDATE bot SET id = $1 WHERE id LIKE $2"

	// accessingQNA := "SELECT answer FROM bot WHERE similarity(question, $1) > 0.3 ORDER BY similarity(question, $1) DESC LIMIT 1"

	// queryData = "SELECT * from userdata WHERE email = $1 AND password = $2"
)
func (s *Bot) Put() error {
	_, err := postgres.Db.Exec(queryQNA, s.Question, s.Answer)
	return err
}
func (s *Bot) DeleteData() error {
	_, err := postgres.Db.Exec(deleteQNA, s.Id)
	return err
}
func (s *Bot) Accessing() error {
	
	fmt.Println(s.Question)

	rows, err := postgres.Db.Query(accessingQNA, "%"+s.Question+"%")
	if err != nil {
		return errors.Wrap(err, "error executing query")
	}
	defer rows.Close()

	// Check if any rows are returned
	if !rows.Next() {
		return sql.ErrNoRows
	}

	// Retrieve the answer from the result set
	err = rows.Scan(&s.Answer)
	if err != nil {
		return errors.Wrap(err, "error scanning rows")
	}

	return nil
}

func GetAllData() ([]Bot, error) {
	const queryGetAll = "SELECT * FROM bot ORDER BY id ASC;"
	table, err := postgres.Db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}
	datas := []Bot{}
	for table.Next() {
		var s Bot
		dbErr := table.Scan(&s.Id,&s.Question, &s.Answer)
		if dbErr != nil {
			return nil, dbErr
		}
		datas = append(datas, s)

	}
	table.Close()
	return datas, nil
}

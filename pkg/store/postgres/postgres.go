package postgres

import (
	"database/sql"
	"fmt"
  "github.com/KussainovAnsar/day-1-2-3/new/KussainovAZH-02/pkg/postgres"
)

func ConnectDB(host, port, user, password, dbname string) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

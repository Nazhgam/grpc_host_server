package storage

import (
	"database/sql"
	"first_proto/api_pb_host"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type storage struct {
	db *sql.DB
}

type Storage interface {
	InsertHost(req *api_pb_host.HostReqeust) (int, error)
}

func NewStorage() (Storage, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	s := &storage{db: db}

	if err := s.createTables(); err != nil {
		return nil, err
	}

	return s, nil
}

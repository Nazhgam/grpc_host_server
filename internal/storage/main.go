package storage

import (
	"database/sql"
	"fmt"

	"github.com/Nagzham/grpc_host_server/pkg/api/editor"

	_ "github.com/mattn/go-sqlite3"
)

type storage struct {
	db *sql.DB
}

type Storage interface {
	InsertHost(req *editor.AddQuestionToQuizRequest) (int, error)
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

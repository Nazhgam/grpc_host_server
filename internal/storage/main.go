package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Nagzham/grpc_host_server/pkg/api/editor"

	_ "github.com/lib/pq"
)

type storage struct {
	db *sql.DB
}

type Storage interface {
	InsertQuiz(req *editor.CreateQuizRequest) (int64, error)
	InsertHost(req *editor.AddQuestionToQuizRequest) (int, error)
}

func NewStorage() (Storage, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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

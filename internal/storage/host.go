package storage

import (
	"github.com/Nagzham/grpc_host_server/pkg/api/editor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *storage) InsertHost(req *editor.AddQuestionToQuizRequest) (int, error) {
	var id int
	if err := s.db.QueryRow(insertHost, req.Question,
		req.WrongAnswers[0], req.WrongAnswers[1], req.WrongAnswers[2], req.CorrectAnswer).Scan(&id); err != nil {
		return 0, status.Error(codes.NotFound, "quiz not found")
	}
	return id, nil
}

package server

import (
	"context"
	"fmt"
	"os"

	"github.com/Nagzham/grpc_host_server/internal/storage"
	"github.com/Nagzham/grpc_host_server/pkg/api/editor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostServer struct {
	*editor.UnimplementedQuizServiceServer
	db storage.Storage
}

func NewHostsServer() *HostServer {
	db, err := storage.NewStorage()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return &HostServer{
		db: db,
	}
}

func (h *HostServer) CreateQuiz(ctx context.Context, req *editor.CreateQuizRequest) (*editor.CreateQuizResponse, error) {
	id, err := h.db.InsertQuiz(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if id == 0 {
		return nil, fmt.Errorf("error while insert to host table. inserted id=0")
	}
	return &editor.CreateQuizResponse{QuizId: id}, nil
}
func (HostServer) GetAllQuizzes(context.Context, *editor.GetAllQuizzesRequest) (*editor.GetAllQuizzesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuizzes not implemented")
}
func (h *HostServer) AddQuestionToQuiz(ctx context.Context, req *editor.AddQuestionToQuizRequest) (*editor.AddQuestionToQuizResponse, error) {
	id, err := h.db.InsertHost(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if id == 0 {
		return nil, fmt.Errorf("error while insert to host table. inserted id=0")
	}
	return nil, status.Errorf(codes.Unimplemented, "method AddQuestionToQuiz not implemented")
}
func (HostServer) GetQuizQuestions(context.Context, *editor.GetQuizQuestionsRequest) (*editor.GetQuizQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuizQuestions not implemented")
}
func (HostServer) EditQuestionInQuiz(context.Context, *editor.EditQuestionInQuizRequest) (*editor.EditQuestionInQuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditQuestionInQuiz not implemented")
}

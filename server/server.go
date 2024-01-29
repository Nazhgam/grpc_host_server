package server

import (
	"context"
	"first_proto/api_pb_host"
	"first_proto/storage"
	"fmt"
	"os"
)

type HostServer struct {
	api_pb_host.UnimplementedHostsServer
	db storage.Storage
}

func NewHostsClient() *HostServer {
	db, err := storage.NewStorage()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return &HostServer{
		db: db,
	}
}

func (h *HostServer) AddQuestion(ctx context.Context, req *api_pb_host.HostReqeust) (*api_pb_host.HostResponse, error) {
	id, err := h.db.InsertHost(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if id == 0 {
		return nil, fmt.Errorf("error while insert to host table. inserted id=0")
	}
	return &api_pb_host.HostResponse{Status: int32(id)}, nil
}

func (h *HostServer) mustEmbedUnimplementedHostsServer() {}

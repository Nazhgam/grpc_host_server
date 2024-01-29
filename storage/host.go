package storage

import (
	"first_proto/api_pb_host"
	"fmt"
)

func (s *storage) InsertHost(req *api_pb_host.HostReqeust) (int, error) {
	var id int
	if err := s.db.QueryRow(insertHost, req.Question,
		req.WrongAns[0], req.WrongAns[1], req.WrongAns[2], req.CorrectAns).Scan(&id); err != nil {
		return 0, fmt.Errorf("error while insert InsertHost: %w", err)
	}
	return id, nil
}

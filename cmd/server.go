package main

import (
	"context"
	"strings"
	"sync"

	pb "github.com/amanjots19/grpc-search/proto"
)

type server struct {
	pb.UnimplementedTransacUserServer
	users map[int64]*pb.UserModel
}

func newServerDI() *server {
	return &server{
		users: make(map[int64]*pb.UserModel),
	}
}

func newServerDIProvider() func() *server {
	var mu sync.Mutex
	var n *server
	return func() *server {
		mu.Lock()
		defer mu.Unlock()
		if n == nil {
			n = newServerDI()
		}
		return n
	}
}

func (s *server) AddUser(ctx context.Context, usr *pb.UserModel) (*pb.APIResp, error) {
	if s.userExists(usr) {
		return &pb.APIResp{
			Error:   "User already exists",
			Success: false,
		}, nil
	}
	s.createUser(usr)

	return &pb.APIResp{
		Message: "User added successfully",
		Success: true,
	}, nil
}

func (s *server) userExists(usr *pb.UserModel) bool {
	if _, ok := s.users[usr.ID]; ok {
		return true
	}

	return false
}

func (s *server) createUser(usr *pb.UserModel) {
	s.users[usr.ID] = usr
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var users []*pb.UserModel
	switch criteria := req.Criteria.(type) {
	case *pb.GetUserRequest_Id:
		user, ok := s.users[criteria.Id]
		if !ok {
			return &pb.GetUserResponse{Success: false, Error: "NO USER FOUND"}, nil
		}
		users = append(users, user)

	case *pb.GetUserRequest_IdsRequest:
		for _, id := range criteria.IdsRequest.Ids {
			for _, user := range s.users {
				if user.ID == id {
					users = append(users, user)
				}
			}
		}
	case *pb.GetUserRequest_Search:
		search := criteria.Search
		for _, user := range s.users {
			if (search.Name != "" && strings.Contains(strings.ToLower(user.Name), strings.ToLower(search.Name))) ||
				(search.Email != "" && strings.Contains(strings.ToLower(user.Email), strings.ToLower(search.Email))) ||
				(search.City != "" && strings.Contains(strings.ToLower(user.City), strings.ToLower(search.City))) ||
				(search.Phone != "" && strings.Contains(strings.ToLower(user.Phone), strings.ToLower(search.Phone))) {
				users = append(users, user)
				break
			}
		}
	}

	return &pb.GetUserResponse{Users: users}, nil

}

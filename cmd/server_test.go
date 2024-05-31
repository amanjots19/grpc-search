package main

import (
	"context"
	"testing"

	pb "github.com/amanjots19/grpc-search/proto"
)

// write test cases for AddUser
func TestAddUser(t *testing.T) {
	// Create a new server instance
	s := newServerDI()

	// Create a test user
	user := &pb.UserModel{
		ID:    1,
		Name:  "John Doe",
		Email: "acv@test.com",
		Phone: "123455433",
	}

	// Add the user to the server
	resp, err := s.AddUser(context.Background(), user)
	if err != nil {
		t.Errorf("Failed to add user: %v", err)
	}

	// Verify the response
	if resp.Success != true {
		t.Errorf("Expected success to be true, got false")
	}

	// Verify that the user was added to the server's user map
	if _, ok := s.users[user.ID]; !ok {
		t.Errorf("Expected user to be added to the user map")
	}
}

func TestAddUser_UserExists(t *testing.T) {
	// Create a new server instance
	s := newServerDI()

	// Create a test user
	user := &pb.UserModel{
		ID:    1,
		Name:  "John Doe",
		Email: "acv@test.com",
		Phone: "123455433",
	}

	// Add the user to the server
	s.AddUser(context.Background(), user)

	// Try adding the same user again
	resp, err := s.AddUser(context.Background(), user)
	if err != nil {
		t.Errorf("Failed to add user: %v", err)
	}

	// Verify the response
	if resp.Success != false {
		t.Errorf("Expected success to be false, got true")
	}
}

func TestGetUser(t *testing.T) {
	// Create a new server instance
	s := newServerDI()

	// Create a test user
	user := &pb.UserModel{
		ID:   1,
		Name: "John Doe",
	}

	// Add the user to the server
	s.AddUser(context.Background(), user)

	// Create a Get User request
	req := &pb.GetUserRequest{
		Criteria: &pb.GetUserRequest_Id{
			Id: 1,
		},
	}

	// Get the user from the server
	resp, err := s.GetUser(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to get user: %v", err)
	}

	// Verify the response
	if resp.Users[0] == nil {
		t.Errorf("Expected user to be returned, got nil")
	}

	if resp.Users[0].ID != user.ID {
		t.Errorf("Expected user ID to be %d, got %d", user.ID, resp.Users[0].ID)
	}
}

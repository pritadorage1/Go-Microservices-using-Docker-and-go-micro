// get-employee-service/main.go
package main

import (
	"fmt"

	// Import the generated protobuf code
	"context"

	"github.com/micro/go-micro"
	pb "github.com/mono-repo/get-employee-service/proto/employee"
)

type repository interface {
	Create(*pb.User) (*pb.User, error)
	GetAll() []*pb.User
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	users []*pb.User
}

func (repo *Repository) Create(user *pb.User) (*pb.User, error) {
	updated := append(repo.users, user)
	repo.users = updated
	return user, nil
}

func (repo *Repository) GetAll() []*pb.User {
	return repo.users
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateEmployee(ctx context.Context, req *pb.User, res *pb.CreateUserResponse) error {

	// Save our consignment
	user, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.User = user
	return nil
}

func (s *service) GetEmployees(ctx context.Context, req *pb.GetUsersRequest, res *pb.GetUsersResponse) error {
	users := s.repo.GetAll()
	res.Users = users
	return nil
}

func main() {

	repo := &Repository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("employee"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterEmployeeServiceHandler(srv.Server(), &service{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

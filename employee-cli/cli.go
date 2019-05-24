package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/mono-repo/get-employee-service/proto/employee"

	micro "github.com/micro/go-micro"
)

const (
	address         = "localhost:50051"
	defaultFilename = "employee.json"
)

func parseFile(file string) (*pb.User, error) {
	var user *pb.User
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &user)
	return user, err
}

func main() {
	service := micro.NewService(micro.Name("employee"))
	service.Init()

	client := pb.NewEmployeeServiceClient("employee", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	user, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateEmployee(context.Background(), user)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetEmployees(context.Background(), &pb.GetUsersRequest{})
	if err != nil {
		log.Fatalf("Could not list employees: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
}

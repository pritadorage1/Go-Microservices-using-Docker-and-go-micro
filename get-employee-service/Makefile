build:
	protoc -I. --go_out=plugins=micro:. \
		proto/employee/employee.proto
	
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t get-employee-service .

run: 
	sudo docker run -p 50051:50051 get-employee-service
	sudo docker run -p 50051:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 \
        get-employee-service
FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD employee.json /app/employee.json
ADD employee-cli /app/employee-cli

CMD ["./employee-cli"]
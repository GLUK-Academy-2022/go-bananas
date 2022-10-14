# Let's Go Bananas!
This is a simple REST API written in Go

## Endpoints
These are the current endpoints available

### GET /hello
A JSON string of the project's name. Example:
```
"Go Bananas"
```

## Run with Docker
To run as a containerized application using docker, follow these steps.

### Install Docker
Check if you have Docker installed
```
docker version
```

If Docker isn't installed, follow the Docker documentation: https://docs.docker.com/get-docker/

### Pull Image
In your terminal, run the following command to pull the latest version:
```
docker pull jaykeharrison/go-bananas
```

### Run Imagine Inside Container
To run the image inside a container, run the following command:
```
docker container run -p <your-chosen-port>:8080 jaykeharrison/go-bananas
```

*important* replace `<your-chosen-port>` with the port of your choice

### Access Endpoints
Perform your requests to the available endpoints on the following url:

http://localhost:`<your-chosen-port>`

Where `<your-chosen-port>`is the what you chose when running the image

Example (port = 8080):
```
curl http://localhost:8080/hello
```
returns
```
"Go Bananas"
```

## Run without Docker


### Setup
To begin, you will need to install Go if you haven't already.

You can follow the official Go docs for how to do this (Linux, Mac, and Windows all supported):

https://go.dev/doc/install

Verify your installation via your terminal:
```
go version
```

### Clone
Clone the repository using your method of choice.

### Run the application
Start the application:
```
go run cmd/main/main.go
```

The application and its endpoint will be accessible via http://localhost:8080/hello
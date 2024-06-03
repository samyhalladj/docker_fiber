# Go Fiber HTTP Server with Pi Calculation

This repository contains a simple HTTP server built with Go Fiber, running inside a Docker container. The server includes a route to estimate the value of Pi using a Monte Carlo method, executed via a shell script.

## Project Structure

projet/

    ├── readme.md
    ├── go.sum
    ├── go.mod
    ├── main.go
    ├── docker
            |
            ├── Dockerfile
            ├── docker-compose.yml
            ├── app # Pre-compiled Go binary
            ├── pi.sh # Shell script to calculate Pi
            └── README.md # Project description


## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (only required if you want to compile the binary yourself)

## Getting Started

### 1. Clone the Repository

    ```sh
    git clone https://github.com/yourusername/your-repo-name.git
    cd your-repo-name


### 2. Build and Run the Docker Container

Ensure that the binary app is pre-compiled and present in the project directory. If it's not, you can compile it yourself:

Then, use Docker Compose to build and run the container:

    docker-compose up --build

### 3. Access the Application

The application will be available at http://localhost:8080.

- GET / : Returns "HiHi"
- GET /pi : Executes the shell script to estimate the value of Pi and returns the result

## Detailed Explanation
### Dockerfile

The Dockerfile sets up a lightweight container using the alpine image. It installs bc, copies the pre-compiled Go binary and the pi.sh script into the container, and exposes port 8080.

    # Use a lightweight base image
    FROM alpine:latest

    # Set the working directory
    WORKDIR /app

    # Install bc
    RUN apk add --no-cache bc

    # Copy the pre-compiled binary and the shell script
    COPY app .
    COPY pi.sh .

    # Expose the port the app runs on
    EXPOSE 8080

    # Command to run the app
    CMD ["./app"]

### docker-compose.yml
The docker-compose.yml file defines the service for the Go Fiber application, mapping port 8080 of the host to port 8080 of the container.

    version: '3'
    services:
    go-fiber-app:
        build: .
        ports:
        - "8080:8080"

### pi.sh
The pi.sh script from (https://lipn.univ-paris13.fr/~cerin/pi.sh) calculates an estimation of Pi using the Monte Carlo method. It generates random points and determines how many fall inside a unit circle, then uses this ratio to estimate Pi.

    #!/bin/sh

    START=1
    END=1000

    TOTAL=0
    INSIDE=0

    while [[ $START -le $END ]]
    do
        x=$RANDOM
        x=`echo "$x / 32767" | bc -l`

        y=$RANDOM
        y=`echo $y / 32767 | bc -l`

        xxyy=`echo "($x * $x + $y * $y) <= 1" | bc -l`

        if [[ $xxyy -eq 1 ]]; then
            INSIDE=$((INSIDE+1))
        fi

        TOTAL=$((TOTAL+1))
        START=$((START + 1))
    done

    echo "$INSIDE*4/$TOTAL" | bc -l

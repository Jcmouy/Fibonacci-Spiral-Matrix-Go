# Dockerfile References: https://docs.docker.com/engine/reference/builder/

### STAGE 1: BUILD ###
FROM golang:1.18.0-alpine as builder
#Add the gcc tools.
RUN apk add build-base
# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
# Create app directory
RUN mkdir /app
# Set the Current Working Directory inside the container
WORKDIR /app
ADD . /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Generate Swagger document
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init --parseDependency --parseInternal -g ./cmd/api/main.go -o ./docs
# Generate dependencies by wire
RUN go install github.com/google/wire/cmd/wire@latest && wire ./internal/wired/wire.go
# Build the Go api
RUN go build -o ./fibospiralmatrixgo ./cmd/api

### STAGE 2: RUN ###
FROM golang:1.18.0-alpine
COPY --from=builder /app/fibospiralmatrixgo /go/bin/fibospiralmatrixgo
# Expose port 8080 to the outside world
EXPOSE 8080
# Run the executable
CMD ["fibospiralmatrixgo"]
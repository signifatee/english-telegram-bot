FROM golang:1.21.5-bullseye

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apt-get update && apt-get upgrade -y && \
    apt-get install bash git ssh -y

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ./main ./cmd/bot/main.go

WORKDIR /app/internal/share/migrations
RUN goose postgres "host=172.17.0.1 user=postgres database=postgres password=pFVhbkwyrkf1wrtkGLGwLH4RybPxH25j sslmode=disable" up

WORKDIR /app
# Expose port 8080 to the outside world
EXPOSE 3131

# Run the executable
CMD ["./main"]
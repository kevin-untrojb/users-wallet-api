FROM golang:1.16

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
WORKDIR cmd/api

RUN go build -o main
EXPOSE 8080
CMD ["./main"]
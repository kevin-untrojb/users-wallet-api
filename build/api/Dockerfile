FROM golang:1.16

# Set necessary environmet variables needed for our image


RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV GIN_MODE=release
ENV IGNORE_GO_GET=true
ENV APPLICATION_PACKAGE=./cmd/api
ENV GO_ENVIRONMENT=dev
ENV GO_RACE_DETECTION_DISABLED="true"
ENV PORT=8000

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -a -o main ./cmd/api

EXPOSE $PORT

CMD ./main
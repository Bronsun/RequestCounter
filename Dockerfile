FROM golang:1.17-alpine as builder

ENV CGO_ENABLED=1
ENV GO111MODULE=on

# Maintainer info
LABEL maintainer="Mateusz Broncel"


# Main workdir
WORKDIR /app

# Downloading all go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Change to workdir where is a server
WORKDIR /app/api/server

# Building go project 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/api/server/main .
COPY --from=builder /app/*.env .

EXPOSE ${PORT}


CMD [ "./main" ]


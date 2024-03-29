# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached in later builds if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Second stage of building

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy compiled object and required file from the previous build
COPY --from=builder /app/main  .
COPY --from=builder /app/config/  ./config
# COPY --from=builder /app/log/  ./log
# COPY --from=builder /app/Build/  ./Build

# Expose the port on which the server will run
EXPOSE 8081

# Command for the entry point of execution in the container
CMD ["./main"]
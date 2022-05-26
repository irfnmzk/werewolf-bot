# Build stage
FROM golang:1.18-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /app
WORKDIR /src

# Copy and download dependency using go mod
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the code into the container
COPY ./ ./

# Build the application
RUN go build -o /app .

# runner
FROM alpine:3 AS final

WORKDIR /app

COPY --from=builder /app /app

# Export necessary port
EXPOSE 5000

# Command to run when starting the container
ENTRYPOINT [ "./app" ]
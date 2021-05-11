FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
		CGO_ENABLED=0 \
		GOOS=linux \
		GOARCH=amd64

# Move to working directory /go/src/app
WORKDIR /go/src

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o app .

FROM alpine:3

ARG PORT=3000
ARG	CORS_DOMAIN=*
ENV PORT=$PORT \
		CORS_DOMAIN=$CORS_DOMAIN

# Move to /dist directory as the place for resulting binary folder
WORKDIR /app
COPY --from=builder /go/src .

# Export necessary port
EXPOSE $PORT

# Command to run when starting the container
CMD ["./app"]
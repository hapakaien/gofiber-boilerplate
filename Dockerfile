# Build stage
FROM docker.io/library/golang:1.16-alpine AS build

# Set necessary environmet variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set workdir
WORKDIR /go/src/app

# Copy files to workdir
COPY . /go/src/app

# Build the app
RUN go build -o /go/bin/app

# Production stage
FROM gcr.io/distroless/base-debian10

# Set args
ARG PORT=3000
ARG CORS_ORIGINS=*

# Set env
ENV PORT=$PORT \
    CORS_ORIGINS=$CORS_ORIGINS

# Copy binary from build stage
COPY --from=build /go/bin/app /

# Export necessary port
EXPOSE $PORT

# Command to run when starting the container
CMD ["/app"]

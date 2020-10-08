FROM golang:alpine

ARG PORT=3000
ARG	CORS_DOMAIN=*

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
		CGO_ENABLED=0 \
		GOOS=linux \
		GOARCH=amd64 \
		PORT=$PORT \
		CORS_DOMAIN=$CORS_DOMAIN

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

# Move to /dist directory as the place for resulting binary folder
WORKDIR /go/dist

# Copy binary from build to main folder
RUN cp /go/src/app .

# Export necessary port
EXPOSE $PORT

# Command to run when starting the container
CMD ["/go/dist/app"]
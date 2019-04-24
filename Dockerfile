FROM golang:latest AS build

LABEL maintainer = "Landon Patmore <landon.patmore@gmail.com>"

WORKDIR $GOPATH/src/jobListingsServer

# SRC . -> DEST .
COPY . .

# Download dependencies and create vendor folder to store them
RUN go get -d -v ./...

# Build a statically-linked Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .

# New build phase -- create binary-only image
FROM alpine:latest

# Add support for HTTPS and time zones
RUN apk update && \
    apk upgrade


WORKDIR /root/

# Copy files from previous build container
COPY --from=build /go/src/jobListingsServer/server ./

# Check results
RUN pwd && find .

# Start the application
CMD ["./server"]

# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
##ENV GIN_MODE=release
ENV PORT=8182

ADD . /app
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# RUN go env -w GO111MODULE=on

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./*.go .

# Build
RUN go build -o main
RUN chmod +x main

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE $PORT

# Run
CMD ["./main"]
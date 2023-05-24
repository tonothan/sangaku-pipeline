# syntax=docker/dockerfile:1

FROM golang:1.20

ENV PORT=8182
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -buildvcs=false -o /sangaku
RUN chmod +x /sangaku

EXPOSE $PORT
CMD [ "/sangaku" ]
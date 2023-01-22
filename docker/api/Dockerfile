FROM golang:1.19-alpine

COPY src/api /go/src/api

WORKDIR /go/src/api/

RUN apk upgrade --update && apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest
RUN go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]

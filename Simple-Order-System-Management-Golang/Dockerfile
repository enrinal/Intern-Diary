
# Builder
FROM golang:1.11.4-alpine3.8 as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make && \
    go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/gitlab.warungpintar.co/enrinal/intern-diary/simple-order

# Force the go compiler to use modules 
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /go/src/gitlab.warungpintar.co/enrinal/intern-diary/simple-order /app

CMD /app/simple-order

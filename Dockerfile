FROM golang:1.14-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk update && apk add --no-cache git

RUN go mod download
RUN go build -o bin/workshop cmd/workshop/main.go

EXPOSE 8080

ENTRYPOINT ["bin/workshop"]


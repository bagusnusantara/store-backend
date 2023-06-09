FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD ./src /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go get github.com/gin-gonic/gin

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
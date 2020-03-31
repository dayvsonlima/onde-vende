FROM golang:1.14.1-alpine3.11

WORKDIR /go/src/website

COPY . .

RUN apk add git

RUN go get github.com/gin-gonic/gin && go get -u github.com/cosmtrek/air
RUN go install github.com/gin-gonic/gin

EXPOSE 8080

CMD ["air", "main.go"]

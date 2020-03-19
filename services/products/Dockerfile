FROM golang:1-alpine

WORKDIR /go/src/products

COPY . .

RUN apk add git

RUN go get github.com/gin-gonic/gin
RUN go install github.com/gin-gonic/gin

EXPOSE 8080

CMD ["go", "run", "main.go"]

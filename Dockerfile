FROM golang:1.16

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

CMD ["./image-crawl-demo", "--devtools-ws-url", "ws://chrome:9222"]

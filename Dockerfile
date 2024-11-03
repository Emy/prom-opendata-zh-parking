FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /prom-opendata-zh-parking

EXPOSE 4277

CMD ["/prom-opendata-zh-parking"]
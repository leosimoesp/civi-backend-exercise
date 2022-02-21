FROM golang:1.17.7-alpine3.15

ENV CGO_ENABLED=0

WORKDIR /app
ADD . .
RUN go build -o cartesian-api ./cmd/main.go

ARG PORT

EXPOSE $PORT

CMD ["/app/cartesian-api"]
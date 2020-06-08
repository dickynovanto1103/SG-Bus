FROM golang

WORKDIR /go/src
COPY go.mod .
RUN go mod download

COPY . .
RUN go build main.go
EXPOSE 8082
CMD ["./main"]
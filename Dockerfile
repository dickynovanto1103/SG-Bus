FROM golang

WORKDIR /go/src
COPY go.* ./
RUN go mod download

COPY main.go .
RUN go build main.go

EXPOSE 8082
CMD ["./main"]
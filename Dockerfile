FROM golang:latest

WORKDIR /src

COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download
COPY . /src/
EXPOSE 8080

RUN go build
CMD ["./Db-Generator"]

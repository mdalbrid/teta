FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download -x
RUN go build -o go-app ./cmd/main.go

CMD ["./go-app"]
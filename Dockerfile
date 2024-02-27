FROM golang:alpine

WORKDIR /src

COPY go.mod go.sum ./

RUN apk add git

RUN  go mod download

COPY . /src

RUN go build -o passgen cmd/passgen/main.go

RUN go install

ENTRYPOINT ["./passgen"]
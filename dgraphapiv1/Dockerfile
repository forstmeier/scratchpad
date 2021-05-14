FROM golang:1.15

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN cd cmd/custom && go build .

WORKDIR cmd/custom

CMD ["./custom"]
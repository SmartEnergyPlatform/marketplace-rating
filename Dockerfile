FROM golang:1.11


COPY . /go/src/marketplace-rating
WORKDIR /go/src/marketplace-rating

ENV GO111MODULE=on

RUN go build

EXPOSE 8080

CMD ./marketplace-rating
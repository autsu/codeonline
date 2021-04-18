FROM golang:latest

WORKDIR /code/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
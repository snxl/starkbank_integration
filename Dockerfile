FROM golang:1.21

WORKDIR /go/src/application

# COPY go.mod ./
# RUN go mod download

COPY . .

# RUN go build -o main src/*.go

EXPOSE 8080

CMD [ "tail", "-f", "/dev/null" ]
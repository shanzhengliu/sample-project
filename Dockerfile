FROM golang:1.20.3

RUN mkdir /application/
WORKDIR /application/
COPY . .

RUN go build -o main ./main.go

ENTRYPOINT [ "./main" ]
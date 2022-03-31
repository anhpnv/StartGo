FROM golang:1.16-alpine
WORKDIR /app

COPY * ./
RUN go build -o go-compile hello.go
EXPOSE 8080

CMD [ "./go-compile" ]

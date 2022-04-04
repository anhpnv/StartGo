FROM golang:1.16-alpine
WORKDIR /app

COPY * ./
RUN go build -o go-compile main.go
EXPOSE 8080

CMD [ "./go-compile" ]

FROM golang:latest

WORKDIR /src/app/

COPY . .

RUN go build -o main src/main.go

EXPOSE 8080

CMD [ "./main" ]
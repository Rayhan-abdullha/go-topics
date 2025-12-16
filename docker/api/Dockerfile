FROM ubuntu:24.04

RUN apt update && apt install -y golang-go

WORKDIR /app

COPY ./server.go ./server.go

CMD ["go", "run", "server.go"]
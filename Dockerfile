FROM golang:1.20.3
RUN mkdir /app

COPY go.mod main.go page.html /app
RUN cd /app && go build

ENTRYPOINT /app/hellofly

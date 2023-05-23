FROM golang:1.20.3
RUN mkdir /app

COPY go.mod main.go page.html /app
RUN cd /app && go build
RUN env > /env.txt

ENTRYPOINT /app/hellofly

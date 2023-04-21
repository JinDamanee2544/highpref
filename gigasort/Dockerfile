FROM golang:1.19.8-alpine3.17

WORKDIR /app

COPY . .

RUN 
RUN go mod tidy
RUN go build -o prog
ENV PATH="${PWD}:${PATH}"
RUN echo $PATH
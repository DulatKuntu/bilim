FROM golang:alpine

RUN apk add git
RUN mkdir /app
ADD . /app 
WORKDIR /app

RUN go build
EXPOSE 8080
EXPOSE 4000

CMD ["/app/bilim"]
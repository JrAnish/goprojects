
FROM golang:latest

LABEL maintainer="Anish JR"

RUN mkdir /app

ADD . /app/

WORKDIR /app

COPY . .

RUN go build -o /app/main .

EXPOSE 8000

WORKDIR /app/

RUN pwd

RUN ls

RUN chmod +x ./webservice
CMD ["./webservice"]
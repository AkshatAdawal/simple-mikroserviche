FROM golang:1.10.0-alpine3.7
MAINTAINER Akshat Adawal

ENV SOURCES /go/src/simple-mikroserviche/simple-mikroserviche

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8089
EXPOSE 8089

ENTRYPOINT simple-mikroserviche

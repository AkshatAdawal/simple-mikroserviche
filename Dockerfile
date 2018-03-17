FROM alpine
MAINTAINER Akshat Adawal

COPY simple-mikroserviche /go/simple-mikroserviche
RUN chmod +x /go/simple-mikroserviche

ENV PORT 8089
EXPOSE 8089

ENTRYPOINT /go/simple-mikroserviche

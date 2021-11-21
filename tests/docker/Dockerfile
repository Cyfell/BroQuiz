FROM golang:alpine

RUN apk update \
    && apk add inotify-tools

COPY scripts/ /scripts
COPY tests/docker /docker

ENV DIRECTORY "cmd/"
ENV BINARY "tocomplete"
ENV CGO_ENABLED 0 

VOLUME /code
WORKDIR /code
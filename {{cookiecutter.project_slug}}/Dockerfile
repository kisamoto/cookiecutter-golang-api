FROM golang:1.8-alpine

ENV GLIDE_VERSION 0.12.3

WORKDIR /go/src/app

COPY . .

RUN wget -O glide.tar.gz https://github.com/Masterminds/glide/releases/download/v${GLIDE_VERSION}/glide-v${GLIDE_VERSION}-linux-amd64.tar.gz

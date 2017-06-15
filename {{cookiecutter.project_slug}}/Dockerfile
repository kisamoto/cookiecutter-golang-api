# Builder image
FROM golang:1.8-alpine as builder
WORKDIR /go/src/{{ cookiecutter.repo }}
COPY . .
RUN go build ./bin/api

# Actual minimal deployment image 
FROM alpine:3.6
ENV APP_DIR /opt/app
WORKDIR $APP_DIR
COPY --from=builder /go/src/{{ cookiecutter.repo }}/api .
RUN chmod -R 777 $APP_DIR/api
CMD ["./api"]

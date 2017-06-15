FROM alpine:3.6

ENV APP_DIR /opt/app
ENV APP_USER_ID 10001

WORKDIR $APP_DIR

COPY ./api .

RUN apk --no-cache add ca-certificates \
  && chown $APP_USER_ID -R $APP_DIR \
  && chmod 700 $APP_DIR/api

CMD ["./api"]

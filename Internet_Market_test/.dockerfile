FROM alpine:latest

RUN mkdir /app

COPY /internal/app/app /app

CMD [ "/app/app" ]
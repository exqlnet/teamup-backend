FROM alpine:3.11

COPY ./user-service /

WORKDIR /

CMD ["/user-service", "--app-id=${APP_ID}", "--app-secret=${APP_SECRET}"]

FROM golang:alpine
RUN apk add --no-cache wget bash
COPY ./make_calls.sh /scripts/
CMD ["/scripts/make_calls.sh"]
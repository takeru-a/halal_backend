FROM golang:1.19

WORKDIR /go/src
COPY ./src .
# RUN apk upgrade --update && \
#     apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/rubenv/sql-migrate/...@v1.2.0
RUN go mod tidy

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
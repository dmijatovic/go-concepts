FROM golang:alpine AS build

# install build tools
RUN apk add --update git curl gcc musl-dev

WORKDIR /home/oauth2

COPY . .

RUN go build -o=oauth2-api


# RELEASE image
FROM alpine:3.12

WORKDIR /home/dv4all

COPY --from=build ./oauth2-api .

EXPOSE 5432
EXPOSE 5433

CMD ["/home/dv4all/oauth2-api"]


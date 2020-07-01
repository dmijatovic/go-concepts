FROM golang:alpine3.12 AS build

RUN apk add --update gcc musl-dev

WORKDIR /home/module

COPY . .

RUN go build -o=demo-app

FROM alpine:3.12

WORKDIR /home/dv4all

COPY --from=build /home/module/demo-app /home/dv4all

CMD ["/home/dv4all/demo-app"]
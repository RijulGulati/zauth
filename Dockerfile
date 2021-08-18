FROM golang:1.16-alpine as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal ./internal
COPY third_party ./third_party
COPY main.go .

RUN go build -o /zauth


FROM alpine:latest
WORKDIR /usr/local/bin
RUN apk add --no-cache libc6-compat
COPY --from=build /zauth .

CMD [ "zauth", "-h" ]


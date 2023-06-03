FROM docker.io/golang:alpine3.17 AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /ifunny-embed

FROM docker.io/alpine:3.17

WORKDIR /

COPY --from=build /ifunny-embed /ifunny-embed

ENV GIN_MODE=release
ENV PORT=6666

EXPOSE 6666

CMD [ "/ifunny-embed" ]

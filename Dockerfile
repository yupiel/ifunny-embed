FROM golang:alpine3.16
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /ifunny-embed

ENV GIN_MODE=release
ENV PORT=6666
EXPOSE 6666

CMD [ "/ifunny-embed" ]
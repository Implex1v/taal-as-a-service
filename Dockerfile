FROM golang:1.17-alpine as build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /app/taas

FROM alpine:3
WORKDIR /app
COPY --from=build /app/taas /app/taas
COPY taal.txt /app
EXPOSE 8080

CMD [ "/app/taas" ]
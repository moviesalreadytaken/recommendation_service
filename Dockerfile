FROM golang:alpine as builder

WORKDIR /recom-build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o recom cmd/main.go

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /recom

COPY --from=builder /recom-build/recom .

EXPOSE 10003
CMD [ "./recom" ]

FROM golang:1.10.3 as builder

WORKDIR /go/src/github.com/truebluejason/p2e-background

RUN go get github.com/go-sql-driver/mysql

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/p2e-background

FROM alpine:latest

WORKDIR /p2e-background

COPY --from=builder /go/bin/p2e-background .

RUN mkdir config && touch config/production.json && chmod +x p2e-background

EXPOSE 4000

CMD [ "/p2e-background/p2e-background" ]
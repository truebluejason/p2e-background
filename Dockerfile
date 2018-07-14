FROM golang:1.10.3 as builder

WORKDIR /go/src/github.com/truebluejason/p2e-background

RUN go get github.com/go-sql-driver/mysql

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/p2e-background

FROM alpine:latest

WORKDIR /p2e-background

RUN mkdir scripts && mkdir config && touch config/production.json

COPY --from=builder /go/bin/p2e-background .

COPY --from=builder /go/src/github.com/truebluejason/p2e-background/scripts/start_production.sh scripts

RUN chmod +x p2e-background

EXPOSE 4000

CMD [ "sh", "/p2e-background/scripts/start_production.sh" ]
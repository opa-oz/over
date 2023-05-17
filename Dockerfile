FROM golang:1.20.4-alpine as builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o /app/over ./main.go

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo

ENV TZ Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/over /app/over

CMD ["./over"]
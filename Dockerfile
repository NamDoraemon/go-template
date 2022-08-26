FROM  golang:1.18-alpine as builder

WORKDIR /root
COPY . .

RUN go build
RUN ls -la

FROM alpine:3.9 as runner

RUN apk --no-cache add ca-certificates curl tzdata

WORKDIR /root

COPY --from=builder /root/go-template /root
# for http
EXPOSE 6001

CMD ["./go-template", "start"]

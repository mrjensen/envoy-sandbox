FROM golang:1.8 as builder
WORKDIR /
ADD ./service1/main.go /
RUN CGO_ENABLED=0 GOOS=linux go build -o service /main.go

FROM envoyproxy/envoy-alpine:latest as envoy
RUN apk update && apk add bash
ADD ./start_service.sh /usr/local/bin/start_service.sh
COPY --from=builder /service .
RUN chmod u+x /usr/local/bin/start_service.sh
ENTRYPOINT /usr/local/bin/start_service.sh

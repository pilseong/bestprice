FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o wemakepriceApp ./cmd/api
RUN chmod +x /app/wemakepriceApp

FROM alpine:latest
RUN apk --no-cache add tzdata
WORKDIR /app
COPY --from=builder /app/wemakepriceApp /app
CMD [ "/app/wemakepriceApp" ]
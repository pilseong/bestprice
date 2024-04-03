FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o backend ./cmd/api
RUN chmod +x /app/backend

FROM alpine:latest
RUN apk --no-cache add tzdata
WORKDIR /app
COPY --from=builder /app/backend /app
EXPOSE 5006
ENV PORT 5006
CMD [ "/app/backend" ]
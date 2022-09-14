FROM golang:alpine as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o ./main .

FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
ENTRYPOINT ["./main"]

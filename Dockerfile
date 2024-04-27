FROM golang:1.22.2-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM gcr.io/distroless/base-debian10

# Copy the built binary
COPY --from=builder /app/app /

CMD ["/app"]

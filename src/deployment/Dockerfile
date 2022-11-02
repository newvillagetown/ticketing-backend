FROM golang:1.19-alpine AS builder
WORKDIR /build/
RUN rm -rf ./*
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -ldflags '-s' -installsuffix cgo -o main .

# FINAL STAGE
FROM alpine:3.15 AS runner
WORKDIR /app/
COPY --from=builder /build/main .
RUN echo 'export $(strings /proc/1/environ | grep AWS_CONTAINER_CREDENTIALS_RELATIVE_URI)' >> /root/.profile
ARG VERSION
ENV VER $VERSION
EXPOSE 80
ENTRYPOINT ["./main"]

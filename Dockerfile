FROM golang:bullseye as builder

RUN apt-get update && apt-get install -y \
    gcc \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /leetcode-daily-rank
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags '-linkmode "external" -extldflags "-static"' -o leetcode-daily-rank .

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /leetcode-daily-rank/config/config.yml.example ./config/config.yml
COPY --from=0 /leetcode-daily-rank/leetcode-daily-rank .

CMD ["./leetcode-daily-rank"]

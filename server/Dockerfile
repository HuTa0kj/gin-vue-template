FROM golang:1.23

WORKDIR /build/src

RUN go install github.com/air-verse/air@latest

COPY ./go.mod go.sum ./
RUN go mod download

COPY air.toml .

# 启动 air
CMD ["air"]
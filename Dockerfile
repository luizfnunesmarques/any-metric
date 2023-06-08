# syntax=docker/dockerfile:1

## build
FROM golang:1.20-buster as builder

WORKDIR /app
ENV GO111MODULE=on

COPY . ./

RUN ls
RUN go mod tidy

RUN go build -o /any-metric

## Runner
FROM gcr.io/distroless/base-debian10
WORKDIR /

COPY --from=builder /any-metric /any-metric
COPY metrics/ /metrics/

ENTRYPOINT ["/any-metric"]

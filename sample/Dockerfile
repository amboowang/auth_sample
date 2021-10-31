FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
# RUN go mod download 

COPY *.go .
RUN go build

FROM alpine as run

WORKDIR /app

# COPY templates templates
COPY --from=build /app/sample /app/

RUN chmod +x sample
ENTRYPOINT [ "./sample" ]
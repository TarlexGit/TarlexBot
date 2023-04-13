FROM golang:1.20-buster AS build
ENV GO111MODULE=on
RUN mkdir -p /app
WORKDIR /hash_app
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY . .
COPY TarlexBot .
COPY .env .
# RUN go build -o .
# RUN rm -rf hash_tools/ api/ go.mod go.sum main.go
CMD ["./TarlexBot"]
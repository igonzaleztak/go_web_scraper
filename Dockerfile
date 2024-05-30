FROM golang:1.22.3-alpine as builder

# Install gcc and make
RUN apk update && apk add --no-cache gcc libc-dev make

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY  . .
RUN make build

ENTRYPOINT ["/app/intelygenz_scraper"]
CMD [""]
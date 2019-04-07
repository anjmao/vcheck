FROM golang:1.12 as gobuilder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-X main.Version=$VERSION" -o /app/vcheck

FROM alpine
RUN apk update && apk add ca-certificates
COPY --from=gobuilder /app/vcheck /app/vcheck
WORKDIR /app
ENTRYPOINT ["./vcheck"]

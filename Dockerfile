FROM golang:1.22.0-alpine3.19 as builder

COPY go.mod go.sum /go/src/github.com/SurkovIlya/statistics-app/
WORKDIR /go/src/github.com/SurkovIlya/statistics-app
RUN go mod download
COPY . /go/src/github.com/SurkovIlya/statistics-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/statistics-app github.com/SurkovIlya/statistics-app


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/SurkovIlya/statistics-app/build/statistics-app /usr/bin/statistics-app

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/statistics-app"]

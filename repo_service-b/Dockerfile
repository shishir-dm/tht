FROM golang:1.7.1 AS build-env
COPY src /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/serviceb src/simple-service-b.go

FROM scratch
COPY --from=build-env /go/bin/serviceb /app/serviceb

EXPOSE 8081
CMD ["/app/serviceb"]
FROM golang:1.7.1 AS build-env
COPY src /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/servicea src/simple-service-a.go

FROM scratch
COPY --from=build-env /go/bin/servicea /app/servicea

EXPOSE 8080
CMD ["/app/servicea"]
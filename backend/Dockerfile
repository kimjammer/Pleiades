FROM golang:1.24.1 AS build-stage

WORKDIR /app/service

#Copy go.mod and go.sum
COPY go.mod ./go.sum ./
#Install modules
RUN go mod download
#Copy code
COPY ./*.go ./
#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /pleiades_backend

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /pleiades_backend /pleiades_backend

EXPOSE 8080

USER nonroot:nonroot

CMD ["/pleiades_backend"]
FROM golang AS build-env
WORKDIR /go/src/app
COPY . .
RUN mkdir -p /build
RUN go build -o=/build/cat cat.go

FROM alpine
WORKDIR /app
COPY --from=build-env /build/cat /app/cat
RUN chmod u+x /app/cat

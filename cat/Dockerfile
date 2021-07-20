FROM golang AS build-env
WORKDIR /go/src/app
COPY . .
RUN mkdir -p /build
RUN go build -o=/build/cat cat.go

FROM ubuntu
COPY --from=build-env /build/cat /build/app
RUN chmod u+x /build/app

RUN /bin/bash

FROM golang:alpine AS build-stage
LABEL org.opencontainers.image.source = "https://github.com/lunarxlark/gvc"

WORKDIR /work
ADD . /work
RUN go build -o gvc .


FROM alpine:latest

COPY --from=build-stage /work/gvc /usr/local/bin/gvc
ENTRYPOINT ["/usr/local/bin/gvc"]

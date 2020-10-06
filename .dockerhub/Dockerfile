FROM golang:1.14 as build-env

ARG SOURCE_BRANCH
ARG SOURCE_COMMIT

WORKDIR /go/src/go-bin-template
ADD . /go/src/go-bin-template

RUN go get -d -v ./...

RUN go build -ldflags "-X main.version=$SOURCE_BRANCH -X main.commit=$SOURCE_COMMIT -X main.date=`date -u +%FT%TZ`" -o /go/bin/go-bin-template

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/go-bin-template /
CMD ["/go-bin-template"]
FROM golang:1.17 as build

WORKDIR /go/src
COPY go.mod main.go /go/src/
# RUN go build -a -tags netgo -installsuffix netgo -o /go/bin/
RUN go install -tags netgo -ldflags '-extldflags "-static"'


FROM gcr.io/distroless/static-debian11
LABEL maintainer="aoki"

WORKDIR /app
COPY --from=build /go/bin/dist-design-pattern /app/dist-design-pattern
ENTRYPOINT ["/app/dist-design-pattern"]


# docker build --no-cache -t go-server:latest .
# docker run -it --rm --entrypoint=sh go-server

# https://okzk.hatenablog.com/entry/2016/08/03/234738
# https://medium.com/veltra-engineering/distroless-e7f877e415c2
# https://github.com/golang-standards/project-layout/blob/master/README_ja.md

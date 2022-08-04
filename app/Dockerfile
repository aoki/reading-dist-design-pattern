FROM golang as build
WORKDIR /app
COPY ["main.go", "go.mod", "./"]
RUN go install && go build

FROM gcr.io/distroless/base:debug as app
# FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=build /app/go-server .
EXPOSE 8080
ENTRYPOINT [ "./go-server"]

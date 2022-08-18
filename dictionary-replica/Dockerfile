FROM golang as build
WORKDIR /app
COPY ["main.go", "go.mod", "./"]
RUN go mod download && go mod tidy && go install && go build

FROM gcr.io/distroless/base:debug as app
# FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=build /app/dictionary-replica .
EXPOSE 8080
ENTRYPOINT [ "./dictionary-replica"]

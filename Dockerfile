FROM golang:1.22 AS build-stage

WORKDIR /app

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/go_template_project

FROM alpine:3 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/go_template_project /go_template_project

EXPOSE 8000

ENTRYPOINT ["/go_template_project"]
FROM golang:1.20-alpine as build-stage
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /dist

FROM gcr.io/distroless/base-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /dist /dist
EXPOSE 8080
CMD [ "/dist" ]
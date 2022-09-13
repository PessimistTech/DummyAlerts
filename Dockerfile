FROM golang:1.18 AS build
WORKDIR /go/src/app
COPY . .
RUN go build .


FROM golang:1.18
WORKDIR /app
COPY --from=build /go/src/app/DummyAlerts ./app
COPY ./.config .
EXPOSE 8080
ENTRYPOINT ./app


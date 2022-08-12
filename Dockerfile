FROM golang:1.18
WORKDIR /app
COPY ./DummyAlerts app
EXPOSE 8080
ENTRYPOINT ./app


##############################Stage 1#########################################
FROM ubuntu

RUN apt-get update && apt-get install -y golang-go

ENV GO111MODULE=off

COPY . .

RUN CGO_ENABLED=0 go build -o /app .

EXPOSE 3000

ENTRYPOINT ["/app"]

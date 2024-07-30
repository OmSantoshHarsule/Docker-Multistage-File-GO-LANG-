##############################Stage 1#########################################
FROM ubuntu AS bulider

RUN apt-get update && apt-get install -y golang-go

ENV GO111MODULE=off

COPY . .

RUN CGO_ENABLED=0 go build -o /app .

###############################Stage 2#########################################
FROM scratch

COPY --from=bulider /app /app

ENTRYPOINT [ "/app" ]

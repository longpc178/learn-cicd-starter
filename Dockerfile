FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

COPY .env ./.env
COPY notely ./notely

CMD ["./notely"]

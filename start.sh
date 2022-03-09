#!/bin/bash

source .env

docker-compose down
docker image prune -f
docker volume prune -f
docker-compose up

# docker exec -it kafka kafka-topics --create --topic $TOPIC --partitions 10 --replication-factor 1 --if-not-exists --bootstrap-server kafka:9092
# docker exec -it kafka kafka-topics --bootstrap-server kafka:9092 --describe --topic cookie-data
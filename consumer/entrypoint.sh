#!/bin/sh

rm -f /opt/consumer/consumer.bin

echo "Building consumer..."
env CGO_ENABLED=0 go build -a -installsuffix cgo -o /opt/consumer/consumer.bin main.go

chmod +x /opt/consumer/consumer.bin

echo "Starting consumer..."
exec /opt/consumer/consumer.bin

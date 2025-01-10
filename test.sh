#!/bin/bash

set -e

echo "Building Go application..."
go build -o ./bin/server ./cmd/server

echo "Starting the Go server..."
./bin/server &
SERVER_PID=$!

sleep 1

echo "Testing the server response..."

HOST="localhost"
PORT=6379
MESSAGE="PING"

# send_request() {
#   RESPONSE=$(echo -ne "$MESSAGE" | nc -w 1 $HOST $PORT | tr -d '\r\n')
#   echo "RESPONSE: $RESPONSE"
# }

# send_request & send_request wait
RESPONSE=$(echo -ne "$MESSAGE" | nc -w 1 $HOST $PORT | tr -d '\r\n')


# echo "All requests completed!"

# Uncomment the following block to validate the response and shutdown the server
if [ "$RESPONSE" == "+PONG" ]; then
  echo "Test passed: Received expected response '+PONG'"
else
  echo "Test failed: Unexpected server response"
  echo "Received: $RESPONSE"
  kill $SERVER_PID
  exit 1
fi

echo "Shutting down the server..."
kill $SERVER_PID

echo "All tests ran successfully!"

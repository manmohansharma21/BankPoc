#!/bin/bash

# Define the port you want to release
PORT_TO_RELEASE=3000

# Find the process using the specified port
PID=$(lsof -t -i :$PORT_TO_RELEASE)

# Check if a process was found
if [ -z "$PID" ]; then
  echo "No process found using port $PORT_TO_RELEASE"
else
  # Kill the process
  echo "Killing process $PID using port $PORT_TO_RELEASE"
  kill $PID
fi


##chmod +x release-port.sh
##./release-port.sh


#!/bin/bash

IMAGE_NAME="twortle"
CONTAINER_NAME="twortle-container"
PORT="3000"

echo "Stopping and removing existing container..."
podman stop $CONTAINER_NAME 2>/dev/null || true
podman rm $CONTAINER_NAME 2>/dev/null || true

echo "Removing old image..."
podman rmi $IMAGE_NAME 2>/dev/null || true

echo "Building new image..."
podman build -t $IMAGE_NAME .

echo "Starting new container on port $PORT..."
# -d runs it in detached mode
podman run -d \
  --name $CONTAINER_NAME \
  -p $PORT:$PORT \
  $IMAGE_NAME

echo "Twortle is now running at http://localhost:$PORT"
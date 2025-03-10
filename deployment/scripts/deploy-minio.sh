#!/bin/bash

# Parameters
MINIO_USER=$1
MINIO_PASSWORD=$2
DOMAIN=$3

# Pull the latest MinIO image
docker pull minio/minio:latest

# Stop and remove existing container if it exists
docker stop minio-server || true
docker rm minio-server || true

# Create data directory if it doesn't exist
mkdir -p ~/minio/data
mkdir -p ~/minio/config

# Run MinIO container
docker run -d \
  --name minio-server \
  --restart always \
  -p 9100:9000 \
  -p 9101:9001 \
  -v ~/minio/data:/data \
  -v ~/minio/config:/root/.minio \
  -e "MINIO_ROOT_USER=${MINIO_USER}" \
  -e "MINIO_ROOT_PASSWORD=${MINIO_PASSWORD}" \
  minio/minio:latest server /data --console-address ":9101"

echo "MinIO deployed successfully on ports 9100 (API) and 9101 (Console)"
echo "Access the MinIO Console at http://$DOMAIN:9101"
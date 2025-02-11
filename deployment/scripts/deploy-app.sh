#!/bin/bash

GITHUB_TOKEN=$1
GITHUB_ACTOR=$2
REPO_NAME=$(echo "$3" | tr '[:upper:]' '[:lower:]')
DB_HOST=$4
DB_PORT=$5
DB_USER=$6
DB_PASSWORD=$7
DB_NAME=$8
DB_SSL_MODE=$9

echo "$GITHUB_TOKEN" | docker login ghcr.io -u "$GITHUB_ACTOR" --password-stdin

docker stop real-estate-management || true
docker rm real-estate-management || true

docker pull "ghcr.io/$REPO_NAME:latest"

docker run -d \
  --name real-estate-management \
  --restart always \
  -e DB_HOST="$DB_HOST" \
  -e DB_PORT="$DB_PORT" \
  -e DB_USER="$DB_USER" \
  -e DB_PASSWORD="$DB_PASSWORD" \
  -e DB_NAME="$DB_NAME" \
  -e DB_SSL_MODE="$DB_SSL_MODE" \
  -p 8080:8080 \
  "ghcr.io/$REPO_NAME:latest"
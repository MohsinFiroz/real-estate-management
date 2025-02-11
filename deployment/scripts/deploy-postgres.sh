#!/bin/bash
DB_USER=$1
DB_PASSWORD=$2
DB_NAME=$3

docker pull postgres:latest

docker stop postgres-db || true
docker rm postgres-db || true

docker volume create postgres_data

docker run -d \
  --name postgres-db \
  --restart always \
  -e POSTGRES_USER="$DB_USER" \
  -e POSTGRES_PASSWORD="$DB_PASSWORD" \
  -e POSTGRES_DB="$DB_NAME" \
  -v postgres_data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres:latest
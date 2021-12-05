#!/bin/sh

echo "Creating docker image"
docker build -f Dockerfile -t rodolfobandeira/go-jwt-auth-api:latest .

echo "Pushing image to Dockerhub registry"
docker push rodolfobandeira/go-jwt-auth-api:latest


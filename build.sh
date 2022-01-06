#!/usr/bin/env bash

docker build -f docker/Dockerfile -t ampregistry:5000/log-kit:latest .
docker push ampregistry:5000/log-kit:latest

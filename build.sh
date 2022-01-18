#!/usr/bin/env bash

git pull
docker build -f docker/Dockerfile -t ampregistry:5000/log-kit:latest .
docker push ampregistry:5000/log-kit:latest

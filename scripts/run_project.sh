#!/bin/bash
APP_NAME=provision-envs
docker build --build-arg APP_NAME=$APP_NAME -t provision-envs . -f deploy/Dockerfile
docker run -it --rm --name $APP_NAME $APP_NAME
# to create a container that stays and never finishes
# docker run -d provision-envs tail -f /dev/null
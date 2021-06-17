#!/bin/bash
docker build -t hello-world-go . -f docker/Dockerfile
docker run -it --rm --name go-container hello-world-go
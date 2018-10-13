#!/bin/sh

VERSION=0.0.1
IMAGE=ynishi/mvno

docker build -t ${IMAGE}:${VERSION} . --no-cache
docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest

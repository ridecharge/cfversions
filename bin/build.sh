#!/bin/sh
bin/bump -p -r && \
docker build -t ridecharge/cfversions:latest . && \
docker build -t ridecharge/cfversions:$(cat VERSION) .

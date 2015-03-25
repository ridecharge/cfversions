#!/bin/sh
bin/bump -p -r && \
docker build -t ridecharge/cfversions:$(cat VERSION) . && \
docker tag ridecharge/cfversions:$(cat VERSION) ridecharge/cfversions:latest

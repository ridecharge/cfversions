FROM ubuntu:14.04

# Install git
RUN apt-get update && apt-get install -y git

# Install go
ADD https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz /tmp/go1.4.2.linux-amd64.tar.gz 
RUN tar -C /usr/local -xzf /tmp/go1.4.2.linux-amd64.tar.gz
RUN mkdir -p /opt/go
ENV GOPATH=/opt/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Install App
COPY . $GOPATH/src/github.com/ridecharge/curbformation-version-service
WORKDIR $GOPATH/src/github.com/ridecharge/curbformation-version-service
RUN go get github.com/tools/godep
RUN godep go install
RUN mv $GOPATH/bin/curbformation-version-service /usr/bin

# Cleanup
RUN rm -r /tmp/*
RUN rm -r $GOROOT
RUN rm -r $GOPATH
RUN apt-get purge -y --auto-remove git

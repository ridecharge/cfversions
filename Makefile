CONTAINER=ridecharge/cfversions
VERSION=$(cat VERSION)


all: test build	push

build:  
	bin/bump -p -r && \
	docker build -t $(CONTAINER):latest . && \
	docker build -t $(CONTAINER):$(VERSION) .

push:
	docker push $(CONTAINER)

test:
	godep go test -cover ./... 

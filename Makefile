all: test build	push clean

build:  
	bin/bump -p -r && \
	docker build -t ridecharge/cfversions:latest . && \
	docker build -t ridecharge/cfversions:$(cat VERSION) .

push:
	docker push ridecharge/cfversions

test:
	godep go test -cover ./... 

clean:
	rm versions/coverage.out
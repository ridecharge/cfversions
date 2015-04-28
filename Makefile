DOCKER_REPO?=registry.gocurb.internal:80
CONTAINER=$(DOCKER_REPO)/cfversions


all: build push clean

build:  
	docker build --no-cache -t $(CONTAINER):latest .

push:
	docker push $(CONTAINER)

test:
	godep go test -cover ./... 

clean:
	docker rmi $(CONTAINER)
# Makefile for building and publishing a Docker image

IMAGE_NAME=reply-node-name
TIMESTAMP := $(shell date +%Y%m%d%H%M%S)
IMAGE_TAG=$(TIMESTAMP)
# Default registry
DOCKER_REGISTRY?=zicongmei

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Tag the image for the registry
tag: build
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest

# Push the image to the registry
push: tag
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest

# Build and push in one command
publish: push

# Clean up local images.
clean:
	docker rmi -f $(IMAGE_NAME):$(IMAGE_TAG) $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) $(DOCKER_REGISTRY)/$(IMAGE_NAME):latest

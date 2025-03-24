# Makefile for building and publishing a Docker image

IMAGE_NAME=echo-server
IMAGE_TAG=latest
# Default registry
DOCKER_REGISTRY?=zicongmei

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Tag the image for the registry
tag: build
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# Push the image to the registry
push: tag
	docker push $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# Build and push in one command
publish: push

# Clean up local images.
clean:
	docker rmi -f $(IMAGE_NAME):$(IMAGE_TAG) $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

include build-spec.env

IMAGE_TAG=$(IMAGE_NAME):$(TAG)
DOCKERFILE=$(DOCKERFILE_PATH)

build:
	docker build -t $(IMAGE_TAG) -f $(DOCKERFILE) .

push: build
	docker push $(IMAGE_TAG)

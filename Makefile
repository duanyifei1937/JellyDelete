DOCKER_HUB = hub.xxx.com
DOCKER_TAG = v0.1
#DOCKER_IMAGE = $(DOCKER_HUB)/library/go-gin-blog
DOCKER_IMAGE = jellydelete

build:
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

push:
	@docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

all: build
# all: build push

rm:
	@docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG)
.PHONY: build push all rm%
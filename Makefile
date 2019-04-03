VERSION:=$(shell git describe --tags --always --dirty)
IMAGE:=anjmao/vcheck:$(VERSION)

container:
	@docker build -t $(IMAGE) .

push:
	@docker push $(IMAGE)
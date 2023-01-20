.PHONY: all build

export DGRAPH_VERSION ?= v22.0.2
export STRUCTURE_KG_VERSON ?= v1.0.0

all: build push

build-arm:
	docker build -f ./Dockerfile.arm64 -t structureorg/kg:${STRUCTURE_KG_VERSON}-arm64 -t structureorg/kg:latest-arm64 --build-arg DGRAPH_VERSION=${DGRAPH_VERSION} .

build-amd:
	docker build -f ./Dockerfile.amd64 -t structureorg/kg:${STRUCTURE_KG_VERSON}-amd64 -t structureorg/kg:latest-amd64 --build-arg DGRAPH_VERSION=${DGRAPH_VERSION} .

build: build-arm build-amd

push:
	@echo "If the following fails, run 'docker login -u <username>'"
	docker manifest create structureorg/kg:${STRUCTURE_KG_VERSON} --amend structureorg/kg:${STRUCTURE_KG_VERSON}-arm64 --amend structureorg/kg:${STRUCTURE_KG_VERSON}-amd64
	docker manifest push structureorg/kg:${STRUCTURE_KG_VERSON}
	docker manifest create structureorg/kg:latest --amend structureorg/kg:latest-arm64 --amend structureorg/kg:latest-amd64
	docker manifest push structureorg/kg:latest

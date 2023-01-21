.PHONY: all build

export DGRAPH_VERSION ?= v22.0.2
export STRUCTURE_KG_VERSON ?= $(shell git tag | sort -V | tail -1)

all: build push

build-arm:
	docker build -f ./Dockerfile.arm64 -t structureorg/kg:${STRUCTURE_KG_VERSON}-arm64 --build-arg DGRAPH_VERSION=${DGRAPH_VERSION} .
	docker tag structureorg/kg:${STRUCTURE_KG_VERSON}-arm64 structureorg/kg:latest-arm64
	docker push structureorg/kg:${STRUCTURE_KG_VERSON}-arm64

build-amd:
	docker build -f ./Dockerfile.amd64 -t structureorg/kg:${STRUCTURE_KG_VERSON}-amd64 --build-arg DGRAPH_VERSION=${DGRAPH_VERSION} .
	docker tag structureorg/kg:${STRUCTURE_KG_VERSON}-amd64 structureorg/kg:latest-amd64
	docker push structureorg/kg:${STRUCTURE_KG_VERSON}-amd64

build: build-arm build-amd

push:
	@echo "If the following fails, run 'docker login -u <username>'"
	docker manifest create structureorg/kg:${STRUCTURE_KG_VERSON} --amend structureorg/kg:${STRUCTURE_KG_VERSON}-arm64 --amend structureorg/kg:${STRUCTURE_KG_VERSON}-amd64
	docker manifest push structureorg/kg:${STRUCTURE_KG_VERSON}
	docker manifest rm structureorg/kg:latest
	docker manifest create structureorg/kg:latest --amend structureorg/kg:${STRUCTURE_KG_VERSON}-arm64 --amend structureorg/kg:${STRUCTURE_KG_VERSON}-amd64
	docker manifest push structureorg/kg:latest

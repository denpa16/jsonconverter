LOCAL_BIN:=$(CURDIR)/bin

DOCKER_DIR=${CURDIR}/build/dev
DOCKER_YML=${DOCKER_DIR}/docker-compose.yml
ENV_NAME="jsonconverter"

# include ${DOCKER_DIR}/.env

include .env

.PHONY: .build
.build:
	go build -o bin/jsonconverter . 

run:
	go run .


test:
	go test ./...

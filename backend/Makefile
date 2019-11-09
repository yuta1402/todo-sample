all: run

init: docker/build docker/up deps

deps:
	dep ensure

up: docker/up

down: docker/down

sh: docker/sh

run:
	go run main.go

docker/build: docker-compose.yml
	docker-compose build

docker/up:
	docker-compose up -d

docker/down:
	docker-compose down

docker/sh:
	docker-compose exec db /bin/bash

.PHONY: init up down sh build run
.PHONY: docker/*

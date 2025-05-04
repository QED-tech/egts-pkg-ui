up: compose-up

compose-up:
	docker compose -f ./infra/docker-compose.yml up --remove-orphans --build -d

#production builds
IMAGE_TAG=$(shell git log | head -n 1 | tail -c 7)

build-gw:
	docker build --platform linux/amd64 --file=infra/production/gw/Dockerfile --tag=qedtechdocker/egts-gw:${IMAGE_TAG} ./infra/production/gw

push-gw:
	docker push qedtechdocker/egts-gw:${IMAGE_TAG}

deploy:
	cd infra/ansible && make deploy

build: build-gw
push: push-gw
ci: build push deploy

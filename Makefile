up: compose-up

compose-up:
	docker compose -f ./infra/docker-compose.yml up --remove-orphans --build -d

#production builds
IMAGE_TAG=$(shell git log | head -n 1 | tail -c 7)

build-front-app:
	cd frontend && npm run build

build-docker-gw:
	docker build --platform linux/amd64 --file=infra/production/gw/Dockerfile --tag=qedtechdocker/egts-gw:${IMAGE_TAG} .

build-gw: build-front-app build-docker-gw

build-migrator:
	docker build --platform linux/amd64 --file=infra/production/migrator/Dockerfile --tag=qedtechdocker/egts-migrator:${IMAGE_TAG} .

build-backend:
	docker build --platform linux/amd64 --file=infra/production/backend/Dockerfile --tag=qedtechdocker/egts-backend:${IMAGE_TAG} ./backend/

push-backend:
	docker push qedtechdocker/egts-backend:${IMAGE_TAG}

push-migrator:
	docker push qedtechdocker/egts-migrator:${IMAGE_TAG}

push-gw:
	docker push qedtechdocker/egts-gw:${IMAGE_TAG}

deploy:
	cd infra/ansible && make deploy

build: build-gw build-backend build-migrator
push: push-gw push-backend push-migrator
ci: build push deploy

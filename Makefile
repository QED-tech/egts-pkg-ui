up: compose-up

compose-up:
	docker compose -f ./infra/docker-compose.yml up --remove-orphans --build -d
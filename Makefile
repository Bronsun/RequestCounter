build-docker:
	docker compose up --build

run-docker:
	docker compose up

clear-docker:
	docker-compose down -v 
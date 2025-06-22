IMAGE_NAME = "sms-service"

build:
	docker build -t ${IMAGE_NAME}:latest .

down:
	docker compose down

up:
	docker compose up

restart:
	docker compose restart
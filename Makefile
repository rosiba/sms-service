IMAGE_NAME = "sms-service"

build:
	docker build -t ${IMAGE_NAME}:latest .

down:
	docker compose down

up:
	docker compose up

restart:
	docker compose restart

rm:
	docker compose rm -fsv
	docker image rm -f sms-service

all: rm build up
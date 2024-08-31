.PHONY: up down setup-env prune

up: setup-env
	docker compose -f hack/development/docker-compose.yml up --build

down:
	docker compose -f hack/development/docker-compose.yml down
	kind delete cluster --name choregate

setup-env:
	hack/development/setup-env.sh

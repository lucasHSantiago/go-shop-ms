# ==============================================================================
# Docker Compose

compose = cd ./config/compose/ && docker compose -f docker-compose.yaml -p compose up -d

compose-up:
	${compose}

compose-build-up:
	${compose} --build

compose-down:
	cd ./config/compose/ && docker compose -f docker-compose.yaml down

compose-logs:
	cd ./config/compose/ && docker compose -f docker-compose.yaml logs

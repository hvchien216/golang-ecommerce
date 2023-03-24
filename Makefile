ifndef PROJECT_NAME
PROJECT_NAME := ecommerce
endif

ifndef DOCKER_BIN:
DOCKER_BIN := docker
endif

ifndef DOCKER_COMPOSE_BIN:
DOCKER_COMPOSE_BIN := docker-compose
endif

COMPOSE := PROJECT_NAME=${PROJECT_NAME} ${DOCKER_COMPOSE_BIN} -f build/docker-compose.yaml
COMPOSE := ${COMPOSE} -f build/docker-compose.dev.yaml

API_COMPOSE = ${COMPOSE} run --rm --service-ports -w /api api

# ----------------------------
# Project level Methods
# ----------------------------
setup: api-setup
boilerplate: api-boilerplate
run: api-run
test: api-test

# ----------------------------
# api Methods
# ----------------------------
build-local-image:
	${DOCKER_BIN} build -f build/local.go.Dockerfile -t ${PROJECT_NAME}-golang-local:latest .
	-${DOCKER_BIN} images -q -f "dangling=true" | xargs ${DOCKER_BIN} rmi -f

api-boilerplate: api-setup api-gen-models api-go-generate
api-setup: pg sleep api-dbup
	${DOCKER_BIN} image inspect ${PROJECT_NAME}-golang-local:latest >/dev/null 2>&1 || make build-local-image
#create the empty migration file
api-cm:
	${COMPOSE} run --rm db-migrate sh -c './migrate create -ext sql -dir /api-migrations -seq $(f)'
api-dbup:
	${COMPOSE} run --rm db-migrate sh -c './migrate -path /api-migrations -database $$DB_URL -verbose up'
api-dbdrop:
	${COMPOSE} run --rm db-migrate sh -c './migrate -path /api-migrations -database $$DB_URL -verbose down'
api-gen-models:
	@${API_COMPOSE} sh -c 'sqlboiler --wipe psql && GOFLAGS="-mod=vendor"'
	@echo "----------sqlboiler is done----------"
gql:
	@cd internal/gql && go run github.com/99designs/gqlgen generate
	@echo "----------gqlgen is done----------"
api-go-generate:
	${API_COMPOSE} sh -c "go generate ./..."
api-test:
	@${API_COMPOSE} sh -c "go test -mod=vendor -failfast -timeout 5m ./..."
#	@${API_COMPOSE} sh -c "go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./..."
api-run:
	@#${API_COMPOSE} sh -c ""
	@${API_COMPOSE} sh -c "go run -mod=vendor cmd/main.go"
api-update-vendor:
	@${API_COMPOSE} sh -c "go mod tidy && go mod vendor"
api-gen-mocks:
	@${API_COMPOSE} sh -c "mockery --dir internal/modules --all --recursive --inpackage"

# ----------------------------
# Base Methods
# ----------------------------
pg:
	${COMPOSE} up -d pg

redis:
	${COMPOSE} up -d redis

sleep:
	sleep 5

.PHONY: boilerplate run test api-run api-test api-setup api-dbmigrate api-dbdrop build-local-go-image pg redis dropdb sleep gql

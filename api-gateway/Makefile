CURRENT_DIR=$(shell pwd)
DBURL=postgres://postgres:root@localhost:5432/newdb?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database '$(DBURL)' -verbose up

mig-down:
	migrate -path migrations -database '$(DBURL)' -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table
# Makefile

SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := internal/http/docs
SWAGGER_INIT := $(SWAGGER) init -g ./internal/http/router.go -o $(SWAGGER_DOCS)

# Target to generate swagger documentation
swag-gen:
	$(SWAGGER_INIT)


run:
	go mod tidy
	clear
	go run cmd/main.go




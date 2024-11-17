
# proto-gen:
# 	protoc --go_out=./ --go-grpc_out=./ gmailapp-submodule/*.proto
	

# migrate_up:
# 	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose up

# migrate_down:
# 	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose down

# migrate_force:
# 	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose force 1

# migrate_file:
# 	migrate create -ext sql -dir internal/migrations -seq create_table

# make spush:
# 	cd gmailapp-submodule && git add . && git commit -m "protos changed" && git push
# 	make proto-gen


CURRENT_DIR=$(shell pwd)
DBURL=postgres://postgres:1234@localhost:5432/tender?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path internal/migrations -database '$(DBURL)' -verbose up

mig-down:
	migrate -path internal/migrations -database '$(DBURL)' -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

mig-force:
	migrate -path ./migrations -database '$(DBURL)' force 1

run:
	go mod tidy
	clear
	go run cmd/main.go




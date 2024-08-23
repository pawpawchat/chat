PROTOC = protoc
PROTO_PATH = api/proto
PROTO_FILES_MASK = ${PROTO_PATH}/*.proto


proto:
	@${PROTOC} --proto_path=${PROTO_PATH} --go_out=. ${PROTO_FILES_MASK}
	@${PROTOC} --proto_path=${PROTO_PATH} --go-grpc_out=. ${PROTO_FILES_MASK}
	@echo "the protobuf files have been rebuild"


MIGRATE = migrate
MIGR_DIR = migrations
DB_SOURCE = ${DB_BASE_URL}/$(1)?sslmode=disable
MIGRATE_BODY = ${MIGRATE} -path ${MIGR_DIR} -database $(call DB_SOURCE,$(db)) 


migrate:
ifndef db
	@$(error parameter db is required)
endif
	${MIGRATE_BODY} up


ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif
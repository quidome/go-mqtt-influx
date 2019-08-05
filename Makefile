PROJECT := mqtt-influx
PKG := github.com/quidome/$(PROJECT)
SRV = mqtt-influx

# Expose compile-time information as linker symbols
LDFLAGS += -X '$(PKG)/diagnostic.AppVersion=${BUILD_VERSION}'
LDFLAGS += -X '$(PKG)/diagnostic.BuildTimestamp=`date +%Y-%m-%dT%H:%M:%S%:z`'
LDFLAGS += -X '$(PKG)/diagnostic.CommitHash=`git rev-parse HEAD`'
LDFLAGS += -X '$(PKG)/diagnostic.GoVersion=`go version`'
# Omit symbol table and debug info, leads to a smaller binary
LDFLAGS += -s

GO_BUILD = CGO_ENABLED=0 go build -a -ldflags "$(LDFLAGS)"

# set default action for make
.DEFAULT_GOAL := help

## Enqueue message into dev environment
dev.send:
	go run cmd/enqueue-message/main.go

## Start development environment and run the server on port 8080
run: dev.start
	echo starting
	CLOUDMQTT_URL='mqtt://localhost:1883/' go run cmd/mqtt-influx/main.go 

## Start the development environment
dev.start:
	# inject commit message validator on start
	# cp -rf .ci-cd/commit-msg .git/hooks
	docker-compose up -d

## Stop the development environment
dev.stop:
	docker-compose down

## Run all unit tests
test:
	go test -race -cover -coverprofile=coverage.out ./...

## Build binary for host arch
build: _build

_build:
	mkdir -p ./.build
	cp Dockerfile .build
	$(GO_BUILD) -o ./.build/$(SRV) $(PKG)/cmd/$(SRV)

## Removes generated build artifacts
clean:
	rm -f ./coverage.out
	rm -f ./r2d2-*
	rm -rf ./.build


dir = ""
i = ""
## Generate a test mock by providing the name of the interface you wish to mock
## make testmock i="InterfaceName"
testmock:
	$(MAKE) _dexec CMD="mockery -dir=$(dir) -name=$(i) -output=$(dir)/internal/mocks"

## Print the help message.
# Parses this Makefile and prints targets that are preceded by "##" comments.
help:
	@echo "" >&2
	@echo "Available targets: " >&2
	@echo "" >&2
	@awk -F : '\
			BEGIN { in_doc = 0; } \
			/^##/ && in_doc == 0 { \
				in_doc = 1; \
				doc_first_line = $$0; \
				sub(/^## */, "", doc_first_line); \
			} \
			$$0 !~ /^#/ && in_doc == 1 { \
				in_doc = 0; \
				if (NF <= 1) { \
					next; \
				} \
				printf "  %-15s %s\n", $$1, doc_first_line; \
			} \
			' <"$(abspath $(lastword $(MAKEFILE_LIST)))" \
		| sort >&2
	@echo "" >&2

_dexec:
	docker exec -it chatbot_dev ${CMD}

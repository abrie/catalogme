PID=hot-compile.pid
PWD=$(shell pwd)
GOPATH=$(PWD)
BIN_DIR=$(PWD)/bin
SRC_DIR=$(PWD)/src/service
BIN_NAME=service
DATA_DIR?=data
PORT?=8400

all: build

initialize-dependencies:
	-@cd $(SRC_DIR); GOPATH=$(GOPATH) go mod init
	-@cd $(SRC_DIR); GOPATH=$(GOPATH) go mod vendor
	-@cd $(SRC_DIR); GOPATH=$(GOPATH) go get ./...

update-dependencies:
	-@cd $(SRC_DIR); GOPATH=$(GOPATH) go get -u ./...

build:
	@cd $(SRC_DIR); GOPATH=$(GOPATH) go build -v -mod=vendor -o $(BIN_DIR)/$(BIN_NAME)

test:
	-@cd $(SRC_DIR); GOPATH=$(GOPATH) go test -coverprofile=$(GOPATH)/coverage.out ./...

test-coverage:
	-@cd $(SRC_DIR); GOPATH=$(GOPATH) go tool cover -html=$(GOPATH)/coverage.out

watch-test: test
	@fswatch -l 0.5 -o $(SRC_DIR) -e ".*" -i "\\.go$$" | xargs -n1 -I{} make test

start: kill build
	@mkdir -p $(DATA_DIR)
	@$(BIN_DIR)/$(BIN_NAME) -d $(DATA_DIR) -p $(PORT) & echo $$! > $(PID)

restart: notify start

notify:
	@echo Recompiling and reloading...

kill:
	-@kill `cat $(PID) 2>/dev/null` 2>/dev/null || true
	-@rm $(PID) || true

watch: restart
	@fswatch -l 0.5 -o $(SRC_DIR) -e ".*" -i "\\.go$$" | xargs -n1 -I{}  make restart

.PHONY: clean
clean:
	-rm $(BIN_DIR)/$(BIN_NAME)

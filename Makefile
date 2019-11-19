PWD=$(shell pwd)
HAPROXY_PORT=$(shell jq .haproxy.port config.json)
STATS_PORT=$(shell jq .stats.port config.json)
FRONTEND_PORT=$(shell jq .frontend.port config.json)
BACKEND_PORT=$(shell jq .backend.port config.json)
DATASTORE_PORT=$(shell jq .datastore.port config.json)

migrate:
	./migration-utils/import.sh
	node ./migration-utils/flatten.js migrated/MERGED.json
	jq '.' migrated/FLAT.json > migrated/data.json

generate-schema:
	node ./migration-utils/generate-schema.js migrated/data.json | jq . > migrated/schema.js

initialize-dependencies:
	@$(MAKE) -C frontend initialize-dependencies
	@$(MAKE) -C services/backend initialize-dependencies
	@echo Dependencies initialized.

up:
	tmux bind-key x source-file tmux/ctrl-c-and-kill.tmux
	tmux bind-key r source-file tmux/kill-and-respawn.tmux
	tmux bind-key a source-file tmux/kill-all.tmux

	tmux split-pane make watch-backend
	tmux set remain-on-exit on
	tmux select-layout even-horizontal

	tmux split-pane make watch-datastore
	tmux set remain-on-exit on
	tmux select-layout even-horizontal

	tmux split-pane make start-haproxy
	tmux set remain-on-exit on
	tmux select-layout even-horizontal

	make watch-frontend

start-haproxy:
	@HAPROXY="127.0.0.1:$(HAPROXY_PORT)" \
	STATS="127.0.0.1:$(STATS_PORT)" \
	FRONTEND="127.0.0.1:$(FRONTEND_PORT)" \
	BACKEND="127.0.0.1:$(BACKEND_PORT)" \
	DATASTORE="127.0.0.1:$(DATASTORE_PORT)" \
	haproxy -f $(PWD)/haproxy/haproxy.cfg

watch-datastore:
	@cd $(PWD)/data; python3 -m http.server $(DATASTORE_PORT)

watch-backend:
	DATA_DIR=$(PWD)/data PORT=$(BACKEND_PORT) $(MAKE) -C services/backend watch

watch-frontend:
	@$(MAKE) -C frontend watch

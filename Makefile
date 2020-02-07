PWD=$(shell pwd)
SOURCE=$(shell jq -r .source secrets/secrets.json)

ORIGINAL=$(PWD)/migrate/original
RESHAPE=$(PWD)/migrate/reshape
GENERATE=$(PWD)/migrate/generate
SERVER=$(PWD)/migrate/server

.PHONY: clean
clean:
	@-rm -rf $(RESHAPE)
	@-rm -rf $(DATASTORE)
	@-rm -rf $(SERVER)

.PHONY: sync
sync:
	INPUT=$(SOURCE) OUTPUT=$(ORIGINAL) ./src/sync/run.sh

.PHONY: reshape
reshape:
	mkdir -p $(RESHAPE)
	./src/reshape/extract-schema.sh $(ORIGINAL)/db.sqlite3 > $(RESHAPE)/original.json
	go run \
		./src/reshape/gen.go \
		$(ORIGINAL)/db.sqlite3 \
		$(RESHAPE)/original.json \
		$(PWD)/src/reshape/reshape.sql.gotmpl \
		> $(RESHAPE)/reshape.sql
	rm -f $(RESHAPE)/db.sqlite3
	sqlite3 $(RESHAPE)/db.sqlite3 < $(RESHAPE)/reshape.sql
	./src/reshape/extract-schema.sh $(RESHAPE)/db.sqlite3 > $(RESHAPE)/shape.json

.PHONY: generate
generate:
	mkdir -p $(GENERATE)
	go run ./src/generate/generate.go \
		-input $(RESHAPE)/shape.json \
		-templates $(PWD)/src/generate \
		-output $(GENERATE)

.PHONY: server
server:
	mkdir -p $(SERVER)
	cp $(GENERATE)/datastore.go $(SERVER)
	cp $(GENERATE)/schema.graphql $(SERVER)
	cp -R ./src/server/* $(SERVER)
	(cd $(SERVER) && go run github.com/99designs/gqlgen -v)

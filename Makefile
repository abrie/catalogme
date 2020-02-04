PWD=$(shell pwd)
SOURCE=$(shell jq -r .source secrets/secrets.json)

.PHONY: clean
clean:
	@-rm -rf $(PWD)/migrate/out

.PHONY: sync
sync:
	INPUT=$(SOURCE) OUTPUT=$(PWD)/migrate/original ./migrate/tools/sync/run.sh

.PHONY: reshape
reshape:
	mkdir -p $(PWD)/migrate/out/reshape
	./migrate/tools/reshape/extract-schema.sh $(PWD)/migrate/original/db.sqlite3 > $(PWD)/migrate/out/reshape/tables.json
	go run \
		./migrate/tools/reshape/gen.go \
		$(PWD)/migrate/original/db.sqlite3 \
		$(PWD)/migrate/out/reshape/tables.json \
		$(PWD)/migrate/tools/reshape/reshape.sql.gotmpl \
		> $(PWD)/migrate/out/reshape/reshape.sql
	rm -f $(PWD)/migrate/out/reshape/db.sqlite3
	sqlite3 $(PWD)/migrate/out/reshape/db.sqlite3 < $(PWD)/migrate/out/reshape/reshape.sql

.PHONY: shape
shape:
	mkdir -p $(PWD)/migrate/out/shape
	./migrate/tools/reshape/extract-schema.sh $(PWD)/migrate/out/reshape/db.sqlite3 > $(PWD)/migrate/out/shape/tables.json


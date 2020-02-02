PWD=$(shell pwd)
SOURCE=$(shell jq -r .source secrets/secrets.json)

.PHONY: clean
clean:
	@-rm -rf $(PWD)/migrate/extracted
	@-rm -rf $(PWD)/migrate/indexed

.PHONY: sync
sync:
	INPUT=$(SOURCE) OUTPUT=$(PWD)/migrate/original ./migration-utils/sync.sh

.PHONY: extract
extract:
	INPUT=$(PWD)/migrate/original OUTPUT=$(PWD)/migrate/extracted ./migration-utils/extract.sh

.PHONY: index
index:
	INPUT=$(PWD)/migrate/extracted OUTPUT=$(PWD)/migrate/indexed ./migration-utils/index.sh

	INPUT=$(PWD)/migrate
.PHONY: merge
merge:
	INPUT=$(PWD)/migrate/extracted OUTPUT=$(PWD)/migrate/merged ./migration-utils/merge.sh

flatten:
	node ./migration-utils/flatten.js $(PWD)/migrate/temp/merged-json/MERGED.json > $(PWD)/migrate/temp/flat-merged-json/FLAT.json
	mkdir -p $(PWD)/migrate/done
	jq '.' $(PWD)/migrate/temp/flat-merged-json/FLAT.json > $(PWD)/migrate/done/data.json

generate-schema:
	node ./migration-utils/generate-schema.js $(PWD)/migrate/done/data.json | jq . > $(PWD)/migrate/done/schema.js

PWD=$(shell pwd)
SOURCE=$(shell jq -r .source secrets/secrets.json)

.PHONY: sync
sync:
	SOURCE=$(SOURCE) ORIGINAL_DATA=$(PWD)/migrate/original ./migration-utils/sync.sh

extract:
	ORIGINAL_DATA=$(PWD)/migrate/original TEMP_DATA=$(PWD)/migrate/temp ./migration-utils/extract.sh

.PHONY: migrate
migrate:
	mkdir -p $(PWD)/migrate/temp/flat-merged-json
	node ./migration-utils/flatten.js $(PWD)/migrate/temp/merged-json/MERGED.json > $(PWD)/migrate/temp/flat-merged-json/FLAT.json
	mkdir -p $(PWD)/migrate/done
	jq '.' $(PWD)/migrate/temp/flat-merged-json/FLAT.json > $(PWD)/migrate/done/data.json

generate-schema:
	node ./migration-utils/generate-schema.js $(PWD)/migrate/done/data.json | jq . > $(PWD)/migrate/done/schema.js

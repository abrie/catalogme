#!/bin/bash

mkdir -p $OUTPUT

for f in $INPUT/*.json
do
  # Get the first one, assume its shape is representative
  jq 'values[keys[0]][0] | keys' $f | sed 's/rowid/ROWID/g' > $OUTPUT/$(basename $f)
done

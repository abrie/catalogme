#!/bin/bash

mkdir -p $OUTPUT

for f in $INPUT/*.json
do
  # Convert the array into a dictionary, indexed by the "rowid" property.
  jq 'values[keys[0]] | map({(.rowid|tostring): .}) | add' $f > $OUTPUT/$(basename $f)
done

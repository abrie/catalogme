#!/bin/bash

mkdir -p $OUTPUT
cat $INPUT/* | jq -s add > $OUTPUT/merged.json

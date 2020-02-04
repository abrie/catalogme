#!/bin/bash

mkdir -p $OUTPUT

TEMP1=$OUTPUT/$RANDOM.json
TEMP2=$OUTPUT/$RANDOM.json
echo "{}" > $TEMP1

for f in $INPUT/*.json
do
  FILENAME=$(basename $f)
  KEY=${FILENAME%.*}
  KEY=$KEY jq '. + {(env.KEY): input}' $TEMP1 $f > $TEMP2
  mv $TEMP2 $TEMP1
done

mv $TEMP1 $OUTPUT/merged.json

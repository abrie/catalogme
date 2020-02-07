#!/bin/bash
set -ue

# This mess extracts the schema from the DB,
# it outputs a JSON dictionary of table names keyed to a descriptive object.
#
# {"table_name":[ {"name":"column1", "type":"string"} ... }

RESULT=tables.json
TEMPB=$RANDOM.tmp
TEMP=$RANDOM.tmp
ACC=$RANDOM.tmp
echo "{}" > $RESULT
for table in $(sqlite3 $1 "select name from sqlite_master where type='table'")
do
  echo "[]" > $ACC
  for column in $(sqlite3 $1 -csv "pragma table_info('$table')" \
    | sed 's/,/ /g' \
    | awk '{printf "{\"name\":\"%s\",\"type\":\"%s\"}\n", $2,$3}')
    do
      echo $column > $TEMPB
      jq '. += [input]' $ACC $TEMPB > $TEMP
      mv $TEMP $ACC
    done
      jq --arg table "$table" '. + {($table):input}' $RESULT $ACC > $TEMP
      mv $TEMP $RESULT
done

cat $RESULT

rm -f $RESULT
rm -f $TEMP
rm -f $TEMPB
rm -f $ACC

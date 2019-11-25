#!/bin/bash
set -ue

mkdir -p $ORIGINAL_DATA

rsync -avp --delete budsbenz.com:/mnt/budsbenz/datastore/ $ORIGINAL_DATA

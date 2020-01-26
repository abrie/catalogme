#!/bin/bash
set -ue

mkdir -p $ORIGINAL_DATA

rsync -avp --delete $SOURCE $ORIGINAL_DATA

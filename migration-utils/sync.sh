#!/bin/bash
set -ue

mkdir -p $OUTPUT

rsync -avp --delete $INPUT $OUTPUT

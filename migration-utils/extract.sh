#!/bin/bash
set -ue

db=$ORIGINAL_DATA/db.sqlite3
extracted=$TEMP_DATA/extracted-json
merged=$TEMP_DATA/merged-json

mkdir -p $extracted
mkdir -p $merged

read -d '' image << EOF || true
SELECT
  json_group_array(
  json_object(
    'rowid', rowid,
    'large_src', large_src,
    'tag', tag,
    'group', "group",
    'href', href,
    'anchor', anchor,
    'sequence', sequence,
    'small_src', small_src))
AS json_result FROM (SELECT rowid, * FROM image)
EOF

read -d '' market_carforsale << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM market_carforsale)
EOF

read -d '' about_topic << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM about_topic)
EOF

read -d '' restoration_item << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM restoration_item)
EOF

read -d '' service_item << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM service_item)
EOF

read -d '' restoration_item << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM restoration_item)
EOF

read -d '' portfolio_group << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM portfolio_group)
EOF

read -d '' portfolio_group_item << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'portfolio_group_id', portfolio_group_id,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM portfolio_group_item)
EOF

read -d '' catalog_series << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM catalog_series)
EOF

read -d '' catalog_series_category << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'catalog_series_id', catalog_series_id,
    'name', name,
    'shortname', shortname,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM catalog_series_category)
EOF

read -d '' catalog_series_category_part << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'catalog_series_category_id', catalog_series_category_id,
    'code', code,
    'internalcode', internalcode,
    'price', price,
    'description', description,
    'tag', tag,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM catalog_series_category_part)
EOF

read -d '' catalog_series_category_part_version << EOF || true
SELECT
	json_group_array(
  json_object(
    'rowid', rowid,
    'catalog_series_category_part_id', catalog_series_category_part_id,
    'code', code,
    'internalcode', internalcode,
    'price', price,
    'description', description,
    'image_group', image_group))
AS json_result FROM (SELECT rowid, * FROM catalog_series_category_part_version)
EOF

sqlite3 $db "$catalog_series" | jq '{"catalog_series":.}' > $extracted/catalog_series.json
sqlite3 $db "$catalog_series_category" | jq '{"catalog_series_category":.}' > $extracted/catalog_series_category.json
sqlite3 $db "$catalog_series_category_part" | jq '{"catalog_series_category_part":.}' > $extracted/catalog_series_category_part.json
sqlite3 $db "$catalog_series_category_part_version" | jq '{"catalog_series_category_part_version":.}' > $extracted/catalog_series_category_part_version.json
sqlite3 $db "$portfolio_group" | jq '{"portfolio_group":.}' > $extracted/portfolio_group.json
sqlite3 $db "$portfolio_group_item" | jq '{"portfolio_group_item":.}' > $extracted/portfolio_group_item.json
sqlite3 $db "$restoration_item" | jq '{"restoration_item":.}' > $extracted/restoration_item.json
sqlite3 $db "$service_item" | jq '{"service_item":.}' > $extracted/service_item.json
sqlite3 $db "$market_carforsale" | jq '{"market_cartforsale":.}' > $extracted/market_carforsale.json
sqlite3 $db "$about_topic" | jq '{"about_topic":.}' > $extracted/about_topic.json
sqlite3 $db "$image" | jq '{"image":.}' > $extracted/image.json

cat $extracted/* | jq -s add > $merged/MERGED.json
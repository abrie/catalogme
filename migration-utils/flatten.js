'use strict';

const fs = require('fs');

const rawdata = fs.readFileSync(process.argv[2]);
const data = JSON.parse(rawdata);
const keys = Object.keys(data);
const flat = keys.reduce( (acc1, key) => {
  acc1[key] = data[key].reduce( (acc2, o) => {
    const id = `${o.rowid}`;
    acc2[id] = o;
    delete(o.rowid);
    o["id"] = id;
    return acc2
  }, {} );
  return acc1;
}, {})

Object.entries(flat).forEach( ([flatName,flatValue]) => {
  if (flatName === "image") return
  Object.entries(flatValue).forEach( ([boId,boObj]) => {
    // Build references and associations
    const references = Object.entries(boObj).filter( ([name,id]) => name.endsWith('_id') );
    const parents = references.forEach( ([name_id,id]) => {
      const parentName = name_id.substr(0, name_id.length - "_id".length);
      const parent = flat[parentName][id];
      if (parent[`list:${flatName}`] === undefined) {
        parent[`list:${flatName}`] = [];
      }
      parent[`list:${flatName}`].push(boId);

    });

    // Add images
    if(boObj.image_group) {
      const images = Object.values(flat.image).filter( (o) => o.group === boObj.image_group);
      const imageIds = images.map( (o) => o.id );
      boObj[`list:image`] = imageIds;
      delete(boObj.image_group);
    } else {
      boObj[`list:image`] = [];
    }
  });
});

const topKeys = [...Object.entries(flat)]
  .filter( ([key,val]) => key.split("_").length === 2 )
  .map( ([key,val]) => {
    const topkey = key.split("_")[0];
    const entry = {};
    entry[`list:${key}`] = [...Object.keys(val)];
    entry[`name`] = `${topkey}`;
    entry[`description`] = `${topkey} description goes here`;
    entry[`shortname`] = `${topkey}`;
    entry[`list:image`] = [];
    return entry;
  });

topKeys.forEach( (t) => flat[t.name] = t );
const str = JSON.stringify(flat);
console.log(str);

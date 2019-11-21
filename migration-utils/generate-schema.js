var data = require('../migrated/data.json')

const printKeys = (object) => {
  return Object.entries(object).map( ([key,value])=> {
    const fieldsObject = Object.entries(value)
      .filter( ([key, value]) => parseInt(key) !== NaN)
      .map( ([key, value]) => value )
      .reduce( (acc, value) => Object.assign(acc, value), {} )
    const fields = Object.entries(fieldsObject)
      .map( ([key,value]) => key === "itemType" ? `list:${value}` : key)
      .map( (key) => key === "images" ? `list:image` : key)
      .filter( (value) => value !== "items" );

    return {key, fields}
  })
}

const splitKey = (key) => key.split("_");
const isLevel2 = (key, prefix) => (splitKey(key).length ===2) && (splitKey(key)[0] === prefix);
const findLevel2 = (prefix, schema) => schema.find( (el) => isLevel2(el.key, prefix) );

const schema = printKeys(data);
const topLevel = [... new Set(schema.map( (f) => splitKey(f.key)[0] ))]
  .filter( (key) => findLevel2(key, schema))
  .map( (key) => {
    const top = findLevel2(key, schema);
    const fields = [
      `list:${top.key}`,
      `name`,
      `shortname`,
      `description`,
      `list:image`,
    ];
    return { key, fields }
  });

console.log(JSON.stringify([...topLevel, ...schema]));

const input = require('./migrated/data.json')

const flattened = Object.entries(input)
  .reduce( (acc, [key,  obj]) =>
    Object.assign(acc, {[key]:[... Object.values(obj)]}), {});

console.log(JSON.stringify(flattened));

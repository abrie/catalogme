<template>
  <div class="page">
    <div>{{name}}</div>
    <h1>{{ params }}</h1>
    <ul>
      <li v-for="field in fields" v-bind:key="field">
        {{field}}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Page',
  props: {
    name: String,
    params: Object,
    fields: Array,
  },
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

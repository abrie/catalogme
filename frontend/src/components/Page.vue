<template>
  <div class="page">
    <h1>{{name}}</h1>
    <p>{{content[name].description}}</p>
    <ul>
      <li v-for="(field,idx) in fields" v-bind:key="idx">
        <p v-if="isListField(field)">
          <ul>
            <li v-for="(obj, idx) in listField(field)" v-bind:key="idx">
              <router-link :to=subFieldLink(obj,field)>{{obj.path}}</router-link>
            </li>
          </ul>
        </p>
        <p v-else>
          {{field}}
        </p>
      </li>
    </ul>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Page',
  props: {
    name: String,
    params: Object,
    fields: Array,
  },
  computed: mapState({
    content: state => state.data.data
  }),
  methods: {
    isListField(field) { return field.startsWith("list:") },
    listField(field) {
      const name = field.split("list:")[1];
      if (name !== "image") {
        return [...Object.values(this.content[name])];
      } else {
        return [];
      }
    },
    subFieldLink(obj, field) {
      const name = field.split("list:")[1];
      const parts = name.split("_");
      parts.reverse().pop();
      parts.reverse();

      const params = {}
      if (obj.path) {
        const p = obj.path.split("/");
        p.reverse().pop();
        p.reduce( (acc, val) => {acc[parts.pop()] = val; return acc}, params );
      }

      return {
        name,
        params,
      }
    }
  }
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

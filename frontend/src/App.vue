<template>
  <div id="app">
    <h1>Catalog Builder</h1>
    <ul>
      <li v-for="link in links" v-bind:key="link">
        <router-link :to=link>{{link}}</router-link>
      </li>
    </ul>
    <hr>
    <router-view></router-view>
    <hr>
  </div>
</template>

<script>

export default {
  name: 'app',
  props: {
    links: Array
  },
  mounted() {
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

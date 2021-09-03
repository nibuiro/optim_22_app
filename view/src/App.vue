<template>
  <div id="app">
    <nav-bar />
    <router-view />
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";

export default {
  name: "App",
  components: {
    "nav-bar": NavBar
  }
};
</script>

<style>
/* Material Design Icons CDN */
@import url("https://cdn.jsdelivr.net/npm/@mdi/font@5.8.55/css/materialdesignicons.min.css");
</style>

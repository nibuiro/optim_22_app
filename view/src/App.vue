<template>
  <div id="app">
    <nav-bar />
    <router-view />
    <copyright />
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
import Copyright from "@/components/Copyright";

export default {
  name: "App",
  components: {
    "nav-bar": NavBar,
    copyright: Copyright
  }
};
</script>

<style>
/* Material Design Icons CDN */
@import url("https://cdn.jsdelivr.net/npm/@mdi/font@5.8.55/css/materialdesignicons.min.css");
</style>

<template>
  <b-navbar>
    <template #brand>
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img
          src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
          alt="Lightweight UI components for Vue.js based on Bulma"
        />
      </b-navbar-item>
    </template>
    <template #start>
      <b-navbar-item href="/">
        Home
      </b-navbar-item>
      <b-navbar-item href="/">
        About
      </b-navbar-item>
    </template>

    <template #end>
      <b-navbar-item tag="div">
        <div class="columns is-variable is-1">
          <register-form class="column" />
          <login-form class="column" />
        </div>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
import RegisterForm from "@/components/RegisterForm";
import LoginForm from "@/components/LoginForm";

export default {
  name: "NavBar",
  components: {
    "register-form": RegisterForm,
    "login-form": LoginForm
  }
};
</script>

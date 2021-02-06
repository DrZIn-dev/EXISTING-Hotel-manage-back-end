<template>
  <div id="app" class="font-monospace patterns">
    <header>
      <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
        <div class="container">
          <a class="navbar-brand" href="/">
            <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="currentColor"
                width="30"
                height="24"
                class="bi bi-journal-check d-inline-block align-top"
                viewBox="0 0 16 16"
            >
              <path
                  fill-rule="evenodd"
                  d="M10.854 6.146a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L7.5 8.793l2.646-2.647a.5.5 0 0 1 .708 0z"
              />
              <path
                  d="M3 0h10a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2v-1h1v1a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v1H1V2a2 2 0 0 1 2-2z"
              />
              <path
                  d="M1 5v-.5a.5.5 0 0 1 1 0V5h.5a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1H1zm0 3v-.5a.5.5 0 0 1 1 0V8h.5a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1H1zm0 3v-.5a.5.5 0 0 1 1 0v.5h.5a.5.5 0 0 1 0 1h-2a.5.5 0 0 1 0-1H1z"
              />
            </svg>
            MURTEL
          </a>
          <form class="d-flex gap-3" v-if="username === ''">
            <button
                class="btn btn-outline-light rounded-0"
                style="text-decoration: none;border-color:#212529"
                @click="$router.push({ name: 'Login' })"
            >
              Login
            </button>
            <button
                class="btn btn-outline-light rounded-0"
                @click="$router.push({ name: 'Register' })"
            >
              Register
            </button>
          </form>
          <form class="d-flex gap-3" v-else>
            <button
                class="btn btn-outline-light rounded-0"
                style="text-decoration: none;border-color:#212529"
                @click="$router.replace({ name: 'Ticket' })"
            >
              WELCOME, {{ username }}
            </button>
            <button
                class="btn btn-outline-secondary rounded-0"
                style="text-decoration: none;border-color:#212529"
                @click="logOut"
            >
              Logout
            </button>
          </form>
        </div>
      </nav>
    </header>
    <div aria-live="polite" aria-atomic="true" class="position-relative">
      <div class="toast-container position-absolute end-0 p-3 mt-5">
        <div class="toast" :class="{'show':isShow}" role="alert" aria-live="assertive" aria-atomic="true">
          <div class="toast-header">
            <strong class="me-auto">{{ title }}</strong>
            <small class="text-muted">just now</small>
            <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
          </div>
          <div class="toast-body">
            {{ msg }}
          </div>
        </div>
      </div>
    </div>
    <div>
      <router-view />
    </div>
  </div>
</template>
<script>
import { call, get } from "vuex-pathify";

export default {
  computed: {
    ...get(["username", "title", "msg", "isShow"])
  },
  methods: {
    ...call(["logOut"])
  }
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  min-height: 100vh;
}
</style>

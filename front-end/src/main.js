import Vue from "vue";
import App from "./App.vue";
import axios from "axios";
import VueAxios from "vue-axios";
import VueCookies from "vue-cookies";
import router from "./router";
import store from "./store";

import "bootstrap";
import "./registerServiceWorker";
import "@/scss/custom.scss";

Vue.config.productionTip = false;

axios.defaults.withCredentials = true;
axios.defaults.baseURL = "http://localhost:3000/api";
Vue.use(VueAxios, axios);
Vue.use(VueCookies);

let app = null;
if (!app) {
  store.dispatch("getUserData").then(() => {
    app = new Vue({
      router,
      store,
      render: h => h(App)
    }).$mount("#app");
  });
}

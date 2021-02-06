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
axios.defaults.baseURL = "https://the-existing-test-back-end.herokuapp.com/api";
Vue.use(VueAxios, axios);
Vue.use(VueCookies);

// let app = null;
// if (!app) {
// store.dispatch("getUserData").then(() => {
//   console.log("reder app");
// });
// }
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");

import Vue from "vue";
import Vuex from "vuex";
import pathify, { updateModules } from "./pathify";
import { make } from "vuex-pathify";

import Register from "@/views/register/store";
import Home from "@/views/home/store";
import Login from "@/views/login/store";
Vue.use(Vuex);

const state = {
    birth_date: "",
    email: "",
    username: ""
};

const store = {
    state,
    getters: {
        ...make.getters(state)
    },
    mutations: {
        ...make.mutations(state)
    },
    actions: {
        ...make.actions(state),
        async getUserData({ commit }) {
            console.log(Vue.$cookies.get("access_token"));
            if (!Vue.$cookies.get("access_token")) {
                await this._vm.axios
                    .get("/user/get-access-token")
                    .then(() => {});
            }
            this._vm.axios.get("/user/private/user").then(({ data }) => {
                const { username } = data;
                commit("username", username);
            });
            console.log("Hello from get User");
        },
        logOut({ commit }) {
            commit("username", "");
            Vue.$cookies.remove("access_token");
            Vue.$cookies.remove("refresh_token");
        }
    },
    modules: {
        Home,
        Register,
        Login
    }
};

const VuexStore = updateModules(
    new Vuex.Store({
        plugins: [pathify.plugin],
        ...store
    })
);

export default VuexStore;

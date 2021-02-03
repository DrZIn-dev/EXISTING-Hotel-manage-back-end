import Vue from "vue";
import Vuex from "vuex";
import pathify, { updateModules } from "./pathify";
import { make } from "vuex-pathify";

import Register from "@/views/register/store";
import Home from "@/views/home/store";
Vue.use(Vuex);

const state = {};
const store = {
    state,
    getters: {
        ...make.getters(state)
    },
    mutations: {
        ...make.mutations(state)
    },
    actions: {
        ...make.actions(state)
    },
    modules: {
        Home,
        Register
    }
};

const VuexStore = updateModules(
    new Vuex.Store({
        plugins: [pathify.plugin],
        ...store
    })
);

export default VuexStore;

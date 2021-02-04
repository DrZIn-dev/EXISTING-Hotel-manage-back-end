import { make } from "vuex-pathify";
import { DateTime } from "luxon";
const getDefaultState = () => ({
    locations: "",
    checkIn: DateTime.local().toISODate(),
    checkOut: DateTime.local()
        .plus({ days: 1 })
        .toISODate(),
    persons: 1
});

const state = getDefaultState();

const getters = {
    ...make.getters(state)
};

const mutations = {
    ...make.mutations(state)
};

const actions = {
    ...make.actions(state)
};

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions
};

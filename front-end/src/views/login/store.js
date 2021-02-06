import { make } from "vuex-pathify";

const getDefaultState = () => ({
    username: "",
    password: ""
});

const state = getDefaultState();

const getters = {
    ...make.getters(state)
};

const mutations = {
    ...make.mutations(state),
    resetState(state) {
        Object.assign(state, getDefaultState());
    }
};

const actions = {
    ...make.actions(state),
    onLogin({ state }) {
        const { username, password } = state;
        console.log({ username, password });
        return this._vm.axios
            .post("/user/signin", {
                identity: username,
                password
            })
            .then(({ data }) => {
                console.log(data);
                return Promise.resolve();
            })
            .catch(err => Promise.reject(err.response.data));
    },
    resetDefaultState({ commit }) {
        commit("resetState");
    }
};

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions
};

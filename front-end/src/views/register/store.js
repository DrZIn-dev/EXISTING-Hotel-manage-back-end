import { make } from "vuex-pathify";

const getDefaultState = () => ({
    username: "",
    password: "",
    email: "",
    birth_date: ""
});

const state = getDefaultState();

const getters = {
    ...make.getters(state)
};

const mutations = {
    ...make.mutations(state)
};

const actions = {
    ...make.actions(state),
    onRegister({ state }) {
        const { username, password, email, birth_date } = state;
        console.log({ username, password, email, birth_date });
        return this._vm.axios
            .post("/user/signup", {
                username,
                password,
                email,
                birth_date
            })
            .then(({ data }) => {
                console.log(data);
                return Promise.resolve();
            })
            .catch(err => {
                console.log(err.response);
                return Promise.reject(err.response.data);
            });
    }
};

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions
};

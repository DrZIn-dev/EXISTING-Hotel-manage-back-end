import { make } from "vuex-pathify";

const getDefaultState = () => ({
  tickets: [],
  hotels: []
});

const state = getDefaultState();

const getters = {
  ...make.getters(state),
  getHotelName(state) {
    return (id) => {
      return state.hotels.find(e => e._id === id);
    };
  }
};

const mutations = {
  ...make.mutations(state)
};

const actions = {
  ...make.actions(state),
  resetDefaultState({ commit }) {
    commit("resetState");
  },
  getTicket({ commit }) {
    this._vm.axios
      .get("/ticket/private/")
      .then(({ data: { tickets, hotels } }) => {
        commit("tickets", tickets);
        commit("hotels", hotels);
        console.log({ tickets, hotels });
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

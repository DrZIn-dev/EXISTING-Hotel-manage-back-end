import { DateTime } from "luxon";
import { make } from "vuex-pathify";

const getDefaultState = () => ({
  name: "",
  checkIn: "",
  checkOut: "",
  persons: 0,
  hotels: []
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
  getHotel({ state, commit }) {
    const { name, checkIn, checkOut } = state;
    console.log(name, checkIn, checkOut);
    this._vm.axios
      .get("/hotel/search", {
        params: { name, cursor: 0, limit: 20, checkIn, checkOut }
      })
      .then(({ data }) => {
        commit("hotels", [...state.hotels, ...data]);
        console.log(data);
      });
  },
  setQueryParams({ commit }, { name, checkIn, checkOut, persons }) {
    commit("name", name);
    commit(
      "checkIn",
      DateTime.fromISO(checkIn).toLocaleString(DateTime.DATE_SHORT)
    );
    commit(
      "checkOut",
      DateTime.fromISO(checkOut).toLocaleString(DateTime.DATE_SHORT)
    );
    commit("persons", persons);
  }
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};

import { make } from "vuex-pathify";

const getDefaultState = () => ({
  id: "",
  name: "",
  checkIn: "",
  checkOut: "",
  persons: 0,
  hotel: {}
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
  createTicket({ state }) {
    const { id, checkIn, checkOut, persons } = state;
    console.log(new Date(checkIn));
    return this._vm.axios
      .post("/ticket/private/", {
        "hotel_id": id,
        "check_in": checkIn,
        "check_out": checkOut,
        "amount": parseInt(persons)
      })
      .then(({ data }) => {
        console.log(data);
        return Promise.resolve();
      })
      .catch(err => Promise.reject(err));
  },
  getHotel({ state, commit }) {
    const { id } = state;
    this._vm.axios
      .get("/hotel", {
        params: { cursor: id }
      })
      .then(({ data }) => {
        commit("hotel", data);
        console.log("hotel"
          , state.hotel
        )
        ;
      });
  },
  setQueryParams({ commit }, { id, name, checkIn, checkOut, persons }) {
    commit("id", id);
    commit("name", name);
    commit(
      "checkIn",
      checkIn
    );
    commit(
      "checkOut",
      checkOut
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

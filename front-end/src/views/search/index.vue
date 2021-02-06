<template>
  <div class=" h-100">
    <div class="container">
      <div class="panel d-flex flex-column align-items-center justify-content-center mx-4">
        <div class="card w-100 my-3 border-0 rounded-0 " v-for="(hotel,index) in hotels" :key="index">
          <div class="row no-gutters">
            <div class="col-lg-2 col-sm-12 ">
              <div class=" h-100 bg-dark" style="min-height: 180px"></div>
            </div>
            <div class="col-lg-10 col-sm-12 d-flex">
              <div class="card-block px-2 py-3 col-10 border-0">
                <h4 class="card-title ">{{ hotel.name }}</h4>
                <p class="card-text mw-100 d-inline-block text-truncate">{{ hotel.description }}</p>
              </div>
              <div class="py-3 pe-3 col-2 d-flex flex-column justify-content-between align-content-end align-items-end">
                <div>
                  <span class="fw-light">฿ </span>
                  <span class="fw-bold h1"> {{ hotel.price }}</span>
                </div>
                <button class="btn btn-success  rounded-0" :disabled="hotel.room < 1" @click="onViewOffer(hotel._id)">
                  View Offers
                </button>
              </div>
            </div>
          </div>

        </div>
        <div class="pb-5">
          <button class="btn btn-dark rounded-0" @click="loadMore">Load More</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { call, get } from "vuex-pathify";

export default {
  computed: {
    ...get("Search", ["hotels", "name",
      "checkIn",
      "checkOut",
      "persons"])
  },
  methods: {
    ...call("Search", ["setQueryParams", "getHotel", "loadMore"]),
    onViewOffer(id) {
      this.$router.replace({
        name: "Offer",
        query: {
          id,
          name: this.locations,
          checkIn: this.checkIn,
          checkOut: this.checkOut,
          persons: this.persons
        }
      });
    }
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.setQueryParams({
        name: vm.$route.query.name,
        checkIn: vm.$route.query.checkIn,
        checkOut: vm.$route.query.checkOut,
        persons: vm.$route.query.persons
      });

      vm.getHotel();
    });
  }
};
</script>

<style>
.panel {
  padding-top: 80px;
}
</style>

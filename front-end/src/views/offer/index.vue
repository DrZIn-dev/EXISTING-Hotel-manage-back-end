<template>
  <div class="h-100 ">
    <div class="container bg-white h-100 ">
      <div class="panel d-flex flex-column  mx-4">
        <div class="d-flex justify-content-between">
          <span class="h1 fw-bold ">{{ hotel.name }}</span>
          <span>
            <span class="fw-light text-black-50">฿ </span>
                  <span class="fw-bold h1 text-dark"> {{ hotel.price }}</span>
          </span>
        </div>
        <div class="w-100 d-flex align-items-end mt-1">
          <button class="ms-auto btn btn-success rounded-0">Book Now</button>
        </div>
        <div class="w-100 row pt-3" style="min-height: 300px">
          <div class="col-lg-6 col-md-12 ">
            <div class="bg-dark w-100 h-100"></div>
          </div>
          <div class="col-lg-3 d-none d-lg-block">
            <div class="bg-dark h-50 mb-1"></div>
            <div class="bg-dark h-50 mt-1"></div>
          </div>
          <div class="col-lg-3 d-none d-lg-block">
            <div class="bg-dark h-50 mb-1"></div>
            <div class="bg-dark h-50 mt-1"></div>
          </div>

        </div>
        <div class="h1 fw-bold mt-5 ">Description</div>
        <div class="mb-5"> {{ hotel.description }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { call, get } from "vuex-pathify";

export default {
  name: "Offers",
  computed: {
    ...get("Offer", ["hotel"])
  },
  methods: {
    ...call("Offer", ["setQueryParams", "getHotel"])
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      const {
        id,
        name,
        checkIn,
        checkOut,
        persons
      } = vm.$route.query;
      vm.setQueryParams({
        id,
        name,
        checkIn,
        checkOut,
        persons
      });
      vm.getHotel();
    });
  }
};
</script>

<style scoped>

</style>
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
          <!-- Button trigger modal -->
          <!-- Button trigger modal -->

          <button type="button" class="ms-auto btn btn-success rounded-0" data-bs-toggle="modal"
                  data-bs-target="#exampleModal">Book Now
          </button>
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


    <!-- Modal -->
    <div class="modal fade rounded-0" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel"
         aria-hidden="true">
      <div class="modal-dialog rounded-0 modal-dialog-centered">
        <div class="modal-content rounded-0">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Confirm Booking</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <table>
              <tr>
                <td>
                  Check In
                </td>
                <td class="fw-bold">
                  : {{ checkIn }}
                </td>
              </tr>
              <tr>
                <td>
                  Check Out
                </td>
                <td class="fw-bold">
                  : {{ checkOut }}
                </td>
              </tr>
              <tr>
                <td>
                  Persons
                </td>
                <td class="fw-bold">
                  : {{ persons }}
                </td>
              </tr>
              <tr>
                <td>
                  Price
                </td>
                <td class="fw-bold ">
                  : {{ hotel.price }} ฿
                </td>
              </tr>
            </table>

          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary rounded-0" data-bs-dismiss="modal">Close</button>
            <button type="button" class="btn btn-success rounded-0">Confirm</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { call, get } from "vuex-pathify";

export default {
  name: "Offers",
  computed: {
    ...get("Offer", {
      hotel: "hotel",
      checkIn: "checkIn",
      checkOut: "checkOut",
      persons: "persons"
    })
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
      console.log(checkIn);
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
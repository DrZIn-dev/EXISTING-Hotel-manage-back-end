<template>
  <div class=" h-100">
    <div class="container">
      <div class="panel mx-4">
        <div class="h1">Tickets</div>
        <div class="row row-cols-4 row-cols-md-4 g-4 pt-4 pb-5">
          <div class="col" v-for="(ticket,index) in tickets" :key="index">
            <div class="card">
              <div class="card-body">
                <h5 class="card-title w-100 d-inline-block text-truncate">{{ hotelName(ticket.hotel_id).name }}</h5>
                <p class="card-text pb-2">
                <table>
                  <tr>
                    <td>Check In</td>
                    <td class="fw-bold">: {{ ticket.check_in.substring(0, 10) }}</td>
                  </tr>
                  <tr>
                    <td>Check Out</td>
                    <td class="fw-bold">: {{ ticket.check_out.substring(0, 10) }}</td>
                  </tr>
                  <tr>
                    <td>Persons</td>
                    <td class="fw-bold">: {{ ticket.amount }}</td>
                  </tr>
                </table>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { call, get } from "vuex-pathify";

export default {
  name: "ticket",
  computed: {
    ...get("Ticket", ["tickets", "getHotelName"])

  },
  methods: {
    ...call("Ticket", ["getTicket", "resetDefaultState"]),
    hotelName(id) {
      console.log(this.getHotelName(id));
      return this.getHotelName(id);
    }
  },
  beforeRouteLeave(to, from, next) {
    this.resetDefaultState();
    next();
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.getTicket();
    });
  }
};
</script>

<style scoped>

</style>
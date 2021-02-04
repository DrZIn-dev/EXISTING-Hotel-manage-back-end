<template>
    <div class=" h-100">
        <div class="container-sm">
            <div class="panel d-flex align-items-center justify-content-center">
                Hello
            </div>
        </div>
    </div>
</template>

<script>
import { call } from "vuex-pathify";
export default {
    computed: {},
    methods: {
        ...call("Search", ["setQueryParams", "getHotel"])
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

<style scoped>
.panel {
    padding-top: 80px;
}
</style>

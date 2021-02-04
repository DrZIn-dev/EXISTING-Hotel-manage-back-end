<template>
    <div class=" h-100">
        <div class="container-sm">
            <div class="panel d-flex align-items-center justify-content-center">
                <div
                    class="card align-self-center rounded-0"
                    style="width: 500px;"
                >
                    <div class="card-body m-2">
                        <h3 class="card-title">Sign In</h3>
                        <h6 class="card-subtitle mb-2 text-muted">
                            For security, please sign in to access your
                            information
                        </h6>
                        <br />

                        <div class="mb-3">
                            <label class="form-label">Username</label>
                            <input
                                type="text"
                                class="form-control rounded-0"
                                placeholder="Please type your name ..."
                                v-model="username"
                            />
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Password</label>
                            <input
                                type="password"
                                class="form-control rounded-0"
                                placeholder="Please type your password ..."
                                v-model="password"
                            />
                        </div>
                        <div>
                            <button
                                class="btn btn-dark w-100 rounded-0"
                                @click="onSubmit"
                                :disabled="isLoading"
                            >
                                <div
                                    class="spinner-grow text-white"
                                    style="width: 1rem; height: 1rem;"
                                    role="status"
                                    v-if="isLoading"
                                ></div>
                                <div v-else>
                                    Sign In
                                </div>
                            </button>
                            <p class="text-danger">{{ err }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { call, sync } from "vuex-pathify";
export default {
    data: () => ({
        isLoading: false,
        err: ""
    }),
    computed: {
        ...sync("Login", ["username", "password"])
    },
    methods: {
        ...call("Login", ["onLogin", "resetDefaultState"]),
        onSubmit() {
            this.isLoading = true;
            this.onLogin()
                .then(() => {
                    this.isLoading = true;
                    this.$router.replace({ name: "Home" });
                })
                .catch(err => {
                    this.err = err;
                });
            console.log("on submit");
        }
    },
    beforeRouteLeave(to, from, next) {
        this.resetDefaultState();
        next();
    }
};
</script>

<style scoped>
.panel {
    padding-top: 80px;
}
</style>

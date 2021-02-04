<template>
    <div class=" h-100">
        <div class="container-sm">
            <div class="panel d-flex align-items-center justify-content-center">
                <div
                    class="card align-self-center rounded-0"
                    style="width: 500px;"
                >
                    <div class="card-body m-2">
                        <h3 class="card-title">Register</h3>
                        <h6 class="card-subtitle mb-2 text-muted">
                            For security, please sign in to access your
                            information
                        </h6>
                        <br />
                        <div class="mb-3">
                            <label class="form-label">Email address</label>
                            <input
                                type="email"
                                class="form-control rounded-0"
                                placeholder="name@example.com"
                                v-model="email"
                            />
                        </div>
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

                        <div class="mb-3">
                            <label class="form-label">Birth Date</label>
                            <input
                                type="date"
                                class="form-control rounded-0"
                                v-model="birth_date"
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
                                    Register
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
import { sync, call } from "vuex-pathify";
export default {
    data: () => ({
        isLoading: false,
        err: ""
    }),
    computed: {
        ...sync("Register", ["username", "password", "email", "birth_date"])
    },
    methods: {
        ...call("Register", ["onRegister"]),
        onSubmit() {
            this.isLoading = true;
            this.onRegister()
                .then(() => {
                    this.$router.replace({ name: "Home" });
                    this.isLoading = false;
                })
                .catch(err => {
                    this.isLoading = false;
                    this.err = err;
                });
        }
    }
};
</script>

<style scoped>
.panel {
    padding-top: 80px;
}
</style>

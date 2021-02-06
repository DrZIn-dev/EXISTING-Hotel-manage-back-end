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
                            <div class="input-group has-validation">
                                <input
                                    type="email"
                                    class="form-control rounded-0"
                                    :class="{
                                        'is-invalid': !validEmail,
                                        'is-valid': validEmail
                                    }"
                                    placeholder="name@example.com"
                                    v-model="email"
                                />

                                <div class="invalid-feedback">
                                    Please choose a e-mail.
                                </div>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Username</label>
                            <div class="input-group has-validation">
                                <input
                                    type="text"
                                    class="form-control rounded-0"
                                    :class="{
                                        'is-invalid': !validUsername,
                                        'is-valid': validUsername
                                    }"
                                    placeholder="Please type your name ..."
                                    v-model="username"
                                    required
                                />
                                <div class="invalid-feedback">
                                    Please choose a username.
                                </div>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Password</label>
                            <div class="input-group has-validation is-invalid">
                                <input
                                    type="password"
                                    class="form-control rounded-0  "
                                    :class="{
                                        'is-invalid': !validPassword,
                                        'is-valid': validPassword
                                    }"
                                    placeholder="Please type your password ..."
                                    v-model="password"
                                />
                                <div class="invalid-feedback">
                                    Length of password should be atleast 8 and
                                    it must be a combination of uppercase
                                    letters, lowercase letters and numbers.
                                </div>
                            </div>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Birth Date</label>
                            <div class="input-group has-validation is-invalid">
                                <input
                                    type="date"
                                    class="form-control rounded-0"
                                    :class="{
                                        'is-invalid': !validBDate,
                                        'is-valid': validBDate
                                    }"
                                    v-model="birth_date"
                                />
                                <div class="invalid-feedback">
                                    Choose a Birth Date.
                                </div>
                            </div>
                        </div>
                        <div>
                            <button
                                class="btn btn-dark w-100 rounded-0"
                                @click="onSubmit"
                                :disabled="isLoading || !validRegister"
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
        ...sync("Register", ["username", "password", "email", "birth_date"]),
        validRegister() {
            return (
                this.validEmail &&
                this.validBDate &&
                this.validUsername &&
                this.validPassword
            );
        },
        validEmail() {
            return this.isEMail(this.email);
        },
        validBDate() {
            console.log("b date", this.birth_date);
            return this.isSpace(this.birth_date) && this.birth_date !== null;
        },
        validUsername() {
            return this.isSpace(this.username);
        },
        validPassword() {
            console.log(
                this.password.length > 8,
                this.isLowerCase(this.password)
            );
            return (
                this.password.length >= 8 &&
                this.isLowerCase(this.password) &&
                this.isUpperCase(this.password) &&
                this.isNumber(this.password)
            );
        }
    },
    methods: {
        ...call("Register", ["onRegister"]),
        isUpperCase(str) {
            return /[A-Z]/.test(str);
        },
        isLowerCase(str) {
            return /[a-z]/.test(str);
        },
        isNumber(str) {
            return /[0-9]/.test(str);
        },
        isSpace(str) {
            return /\S/.test(str);
        },
        isEMail(str) {
            let re = /\S+@\S+\.\S+/;
            return re.test(str);
        },
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

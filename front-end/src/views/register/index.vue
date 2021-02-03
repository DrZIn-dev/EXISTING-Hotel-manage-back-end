<template>
    <div class="patterns h-100">
        <div class="container-sm h-100 ">
            <div class="panel d-flex align-items-center justify-content-center">
                <div class="card align-self-center" style="width: 500px;">
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
                                class="form-control"
                                placeholder="name@example.com"
                            />
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Username</label>
                            <input
                                type="text"
                                class="form-control"
                                placeholder="Please type your name ..."
                            />
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Password</label>
                            <input
                                type="password"
                                class="form-control"
                                placeholder="Please type your password ..."
                            />
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Birth Date</label>
                            <input type="date" class="form-control" />
                        </div>
                        <div>
                            <button class="btn btn-dark w-100">Register</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {};
</script>

<style scoped>
.panel {
    padding-top: 70px;
}
</style>

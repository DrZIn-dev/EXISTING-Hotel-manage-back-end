<template>
    <div class="patterns h-100">
        <div class="container ">
            <div class="hero ">
                <div class="row">
                    <div class="col-lg-12">
                        <div class="hero_text">
                            <h5>WELCOME TO MURTEL</h5>
                            <h2 style="line-height: 82px;">
                                Experience the greatest for you holidays.
                            </h2>
                        </div>
                        <form
                            class="row bg-white position-relative"
                            style="padding: 30px 30px 30px 40px;"
                        >
                            <div class="col-xl-3  col-lg-12">
                                <div>
                                    <label class="pb-3 ">Location</label>
                                    <input
                                        class="form-control rounded-0"
                                        placeholder="Type to search..."
                                    />
                                </div>
                            </div>
                            <div class="col-xl-3  col-lg-12">
                                <div>
                                    <label class="pb-3">Check In</label>
                                    <input
                                        class="form-control rounded-0"
                                        type="date"
                                    />
                                </div>
                            </div>
                            <div class="col-xl-3  col-lg-12">
                                <div>
                                    <label class="pb-3">Check Out</label>
                                    <input
                                        class="form-control rounded-0"
                                        type="date"
                                    />
                                </div>
                            </div>
                            <div class="col-xl-1 col-lg-12">
                                <div>
                                    <label class="pb-3">Person</label>
                                    <input
                                        class="form-control rounded-0"
                                        type="number"
                                        value="1"
                                        min="1"
                                    />
                                </div>
                            </div>
                            <div class="col-xl-2 col-lg-12">
                                <button
                                    class="rounded-0 btn btn-dark w-100 h-100"
                                    type="submit"
                                >
                                    BOOK NOW
                                </button>
                            </div>
                        </form>
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
.hero_text {
    text-align: center;
    max-width: 700px;
    margin: 0 auto 65px;
}
.hero {
    height: 100%;
    padding-top: 145px;
}
.patterns {
    background-color: #dfdbe5;
    background-image: url("data:image/svg+xml,%3Csvg width='80' height='80' viewBox='0 0 80 80' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%239C92AC' fill-opacity='0.4'%3E%3Cpath d='M50 50c0-5.523 4.477-10 10-10s10 4.477 10 10-4.477 10-10 10c0 5.523-4.477 10-10 10s-10-4.477-10-10 4.477-10 10-10zM10 10c0-5.523 4.477-10 10-10s10 4.477 10 10-4.477 10-10 10c0 5.523-4.477 10-10 10S0 25.523 0 20s4.477-10 10-10zm10 8c4.418 0 8-3.582 8-8s-3.582-8-8-8-8 3.582-8 8 3.582 8 8 8zm40 40c4.418 0 8-3.582 8-8s-3.582-8-8-8-8 3.582-8 8 3.582 8 8 8z' /%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
}
</style>

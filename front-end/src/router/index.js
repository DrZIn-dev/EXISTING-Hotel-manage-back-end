import Vue from "vue";
import VueRouter from "vue-router";
import store from "@/store";
import Home from "@/views/home";
import Register from "@/views/register";
import Login from "@/views/login";
import Search from "@/views/search";
import Offer from "@/views/offer";
import Ticket from "@/views/ticket";

Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home
    },
    {
        path: "/register",
        name: "Register",
        component: Register
    },
    {
        path: "/login",
        name: "Login",
        component: Login
    },
    {
        path: "/search",
        name: "Search",
        component: Search
    },
    {
        path: "/offer",
        name: "Offer",
        component: Offer
    },
    {
        path: "/ticket",
        name: "Ticket",
        component: Ticket
    }
];

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes
});

router.beforeEach((to, from, next) => {
    try {
        store.dispatch("getUserData");
    } catch (err) {
        console.log(err);
    }
    next();
});
export default router;

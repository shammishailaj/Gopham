import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/analysis",
      name: "analysis",
      component: () => import("./views/Analysis.vue")
    },
    {
      path: "/details",
      name: "details",
      component: () => import("./views/Details.vue"),
      props: true
    }
  ]
});

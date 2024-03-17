import Vue from "vue";
import Router from "vue-router";
import MainPage from "./components/MainPage.vue";
import AboutPage from "./components/AboutPage.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "Home",
      component: MainPage,
    },
    {
      path: "/about",
      name: "About",
      component: AboutPage,
    },
  ],
  mode: "history",
});

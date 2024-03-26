import Vue from "vue";
import Router from "vue-router";
import MainPage from "./pages/MainPage.vue";
import AboutPage from "./pages/AboutPage.vue";
import ResultsPage from "./pages/ResultsPage.vue";
import ResultDetailsPage from "./pages/ResultDetailsPage.vue";

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
    {
      path: "/results",
      name: "Results",
      component: ResultsPage,
    },
    {
      path: "/results/:id",
      name: "ResultDetails",
      component: ResultDetailsPage,
      props: true,
    },
  ],
  mode: "history",
});

import Vue from "vue";
import Vuex from "vuex";
import resultsModule from "./modules/results";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    results: resultsModule,
  },
});

import axios from "axios";
import { parseText } from "./../../utils/parseText";
import { formatDate } from "./../../utils/formatDate.js";

const resultsModule = {
  state: {
    results: [],
  },
  mutations: {
    setResults(state, results) {
      state.results = results;
    },
    deleteResult(state, resultId) {
      state.results = state.results.filter(
        (result) => result.resultId !== resultId
      );
    },
  },
  actions: {
    async fetchResults({ commit }) {
      try {
        const response = await axios.get("http://localhost:8080/api/results");
        if (!response.data || response.data.length === 0) {
          console.log("No data received or data is empty");
          commit("setResults", []);
          return;
        }
        const results = Array.isArray(response.data)
          ? response.data.map((item, index) => ({
              resultId: item.result_id,
              id: index + 1,
              text: item.text ? parseText(item.text) : "",
              resultName: item.result_name,
              link: item.link,
              keywords: item.keywords,
              date: item.timestamp ? formatDate(item.timestamp) : "",
            }))
          : [];
        commit("setResults", results);
      } catch (error) {
        console.error("Error fetching results:", error);
        throw error;
      }
    },
    async deleteResult({ commit }, resultId) {
      try {
        await axios.delete(`http://localhost:8080/api/results/${resultId}`);
        commit("deleteResult", resultId);
      } catch (error) {
        console.error("Error deleting result:", error);
        throw error;
      }
    },
  },
  getters: {
    getResults(state) {
      return state.results;
    },
  },
};

export default resultsModule;

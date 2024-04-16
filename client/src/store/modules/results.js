import axios from "axios";
import { formatDate } from "./../../utils/formatDate";

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
          commit("setResults", []);
          return;
        }
        const results = Array.isArray(response.data)
          ? response.data
              .sort((a, b) => {
                return new Date(a.timestamp) - new Date(b.timestamp);
              })
              .map((item, index) => ({
                resultId: item.result_id,
                id: index + 1,
                text: item.text,
                resultName: item.result_name,
                link: item.link,
                keywords: item.keywords,
                date: formatDate(item.timestamp),
              }))
          : [];
        commit("setResults", results);
      } catch (error) {
        console.error("Error fetching results: ", error);
        throw error;
      }
    },
    async deleteResult({ commit }, resultId) {
      try {
        await axios.delete(`http://localhost:8080/api/results/${resultId}`);
        commit("deleteResult", resultId);
      } catch (error) {
        console.error("Error deleting result: ", error);
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

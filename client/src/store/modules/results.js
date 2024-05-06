import axios from "axios";
import { formatDate } from "./../../utils/formatDate";
import { parseText } from "./../../utils/parseText";

const baseURL = process.env.VUE_APP_BASE_URL;
const apiVersion = process.env.VUE_APP_API_VERSION;

const resultsModule = {
  state: {
    results: [],
  },
  mutations: {
    submitLink(state, results) {
      state.results.push(results);
    },
    fetchResults(state, results) {
      state.results = results;
    },
    saveResults(state, results) {
      state.results.push(results);
    },
    deleteResult(state, resultId) {
      state.results = state.results.filter(
        (result) => result.resultId !== resultId
      );
    },
  },
  actions: {
    async submitLink({ commit }, payload) {
      try {
        const response = await axios.post(
          `${baseURL}/api/${apiVersion}/scrape`,
          payload
        );
        commit("submitLink", {
          url: payload.url,
          keywords: payload.keywords.map((keyword) => keyword.toLowerCase()),
        });
        return response.data;
      } catch (error) {
        console.error("Error submitting results: ", error);
        throw error;
      }
    },
    async fetchResults({ commit }) {
      try {
        const response = await axios.get(
          `${baseURL}/api/${apiVersion}/results`
        );
        if (!response.data || response.data.length === 0) {
          commit("fetchResults", []);
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
                text: parseText(item.text),
                resultName: item.result_name,
                link: item.link,
                keywords: item.keywords,
                date: formatDate(item.timestamp),
              }))
          : [];
        commit("fetchResults", results);
      } catch (error) {
        console.error("Error fetching results: ", error);
        throw error;
      }
    },
    async saveResults({ commit }, payload) {
      try {
        await axios.post(`${baseURL}/api/${apiVersion}/results`, payload);
        commit("saveResults", {
          url: payload.url,
          keywords: payload.keywords,
          resultName: payload.resultName,
          sentences: payload.sentences,
        });
      } catch (error) {
        console.error("Error saving results: ", error);
        throw error;
      }
    },
    async deleteResult({ commit }, resultId) {
      try {
        await axios.delete(`${baseURL}/api/${apiVersion}/results/${resultId}`);
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

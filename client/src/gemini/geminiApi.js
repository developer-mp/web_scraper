import axios from "axios";

const geminiApi = axios.create({
  baseURL: process.env.VUE_APP_GEMINI_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export function summarizeTextGemini(apiKey, text) {
  const apiEndpoint = process.env.VUE_APP_GEMINI_API_ENDPOINT;
  const data = {
    contents: [
      {
        parts: [
          {
            text: `Summarize the following text: ${text}`,
          },
        ],
      },
    ],
  };
  return geminiApi.post(`${apiEndpoint}${apiKey}`, data);
}

export function analyzeSentimentTextGemini(apiKey, text) {
  const apiEndpoint = process.env.VUE_APP_GEMINI_API_ENDPOINT;
  const data = {
    contents: [
      {
        parts: [
          {
            text: `Perform sentiment analysis of the text in the format Sentiment: [sentiment] and Explanation: [explanation]: ${text}`,
          },
        ],
      },
    ],
  };
  return geminiApi.post(`${apiEndpoint}${apiKey}`, data);
}

export function translateTextGemini(apiKey, language, text) {
  const apiEndpoint = process.env.VUE_APP_GEMINI_API_ENDPOINT;
  const data = {
    contents: [
      {
        parts: [
          {
            text: `Translate the following text into ${language}: ${text}`,
          },
        ],
      },
    ],
  };
  return geminiApi.post(`${apiEndpoint}${apiKey}`, data);
}

export default geminiApi;

import axios from "axios";

const geminiApi = axios.create({
  baseURL: process.env.VUE_APP_GEMINI_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export function generateContent(apiKey, text) {
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

export default geminiApi;

import axios from "axios";

const geminiApi = axios.create({
  baseURL: "https://generativelanguage.googleapis.com/v1beta",
  headers: {
    "Content-Type": "application/json",
  },
});

export function generateContent(apiKey, text) {
  const data = {
    contents: [
      {
        parts: [
          {
            text: text,
          },
        ],
      },
    ],
  };
  return geminiApi.post(
    `/models/gemini-pro:generateContent?key=${apiKey}`,
    data
  );
}

export default geminiApi;

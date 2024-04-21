<template>
  <div>
    <NavBarComponent />
    <div class="container-details">
      <h2 class="details-header">Scraping Result Details</h2>
      <ul class="details-list">
        <li class="details-buttons-group">
          <b-dropdown right text="Action" class="details-button">
            <b-dropdown-item @click="summarizeText"
              >Summarization</b-dropdown-item
            >
            <b-dropdown-item @click="analyzeSentimentText"
              >Sentiment Analysis</b-dropdown-item
            >
            <b-dropdown-item @click="showTranslateModal"
              >Translation</b-dropdown-item
            >
          </b-dropdown>
          <b-button v-b-toggle="'result-statistics'" class="details-button"
            >Statistics</b-button
          >
          <b-collapse id="result-statistics">
            <b-card class="details-card-container"
              ><div v-if="resultStatistics">
                <ul class="details-card">
                  <li>
                    <strong>Word Count:</strong>
                    {{ resultStatistics.wordCount }}
                  </li>
                  <li>
                    <strong>Character Count:</strong>
                    {{ resultStatistics.characterCount }}
                  </li>
                  <li>
                    <strong>Sentence Count:</strong>
                    {{ resultStatistics.sentenceCount }}
                  </li>
                  <div
                    v-for="(count, keyword) in resultStatistics.keywordCounts"
                    :key="keyword"
                  >
                    <li>
                      <strong>Keyword &lt;{{ keyword }}&gt; Count:</strong>
                      {{ count }}
                    </li>
                  </div>
                </ul>
              </div>
              <div v-else>No statistics available</div></b-card
            >
          </b-collapse>
        </li>
        <li><strong>Id:</strong> {{ result.id }}</li>
        <li><strong>Result Name:</strong> {{ result.resultName }}</li>
        <li><strong>Link:</strong> {{ result.link }}</li>
        <li><strong>Keywords:</strong> {{ result.keywords }}</li>
        <li><strong>Text Sample:</strong> {{ result.text }}</li>
        <li><strong>Date:</strong> {{ result.date }}</li>
      </ul>
    </div>
    <b-modal id="translationModal" title="Language Input" hide-footer>
      <b-form @submit.prevent="translateText">
        <label for="languageInput" style="margin-right: 1em"
          >Enter Language:</label
        >
        <input
          type="text"
          v-model="language"
          id="languageInput"
          placeholder="e.g., French, Spanish"
        />
      </b-form>
      <div class="button-container">
        <b-button
          type="submit"
          variant="primary"
          @click="translateText"
          class="button mt-3"
          >Done</b-button
        >
        <b-button
          variant="secondary"
          @click="canceltranslationModal"
          class="button mt-3"
          >Cancel</b-button
        >
      </div>
    </b-modal>
    <ModalComponent
      ref="summarizationModal"
      :modalTitle="summarizedTextTitle"
      :modalText="summarizedText"
      :modalMessage="noSummarizedTextMessage"
      @confirm="downloadSummary"
      @cancel="cancelPreview"
    />
    <ModalComponent
      ref="sentimentAnalysisModal"
      :modalTitle="analyzedSentimentTextTitle"
      :modalText="analyzedSentimentText"
      :modalMessage="noAnalyzedSentimentTextMessage"
      @confirm="downloadSentimentAnalysis"
      @cancel="cancelPreview"
    />
    <FooterComponent />
  </div>
</template>

<style>
.container-details {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.details-header {
  margin-top: 1em;
  margin-bottom: 1em;
}

.details-list {
  list-style-type: none;
}

.details-list li {
  text-align: left;
}

.details-list li strong {
  display: inline-block;
  width: 10em;
}

.details-buttons-group {
  margin-bottom: 1em;
}

.details-button {
  width: 6em;
  margin-right: 0.5em;
}

.details-card {
  list-style-type: none;
}

.details-card li strong {
  display: inline-block;
  width: 15em;
}

.details-card-container * {
  padding: 0;
}

.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 1em;
}
</style>

<script>
import { mapGetters, mapActions } from "vuex";
import FooterComponent from "./../components/FooterComponent.vue";
import NavBarComponent from "./../components/NavBarComponent.vue";
import { cutString } from "./../utils/cutString.js";
import { downloadFile } from "./../utils/downloadFile.js";
import {
  summarizeTextGemini,
  analyzeSentimentTextGemini,
  translateTextGemini,
} from "./../gemini/geminiApi";
import ModalComponent from "./../components/ModalComponent.vue";

export default {
  components: {
    FooterComponent,
    NavBarComponent,
    ModalComponent,
  },
  computed: {
    ...mapGetters(["getResults"]),
    result() {
      const results = this.getResults;
      const result = results.find(
        (result) => result.id === parseInt(this.$route.params.id)
      );
      if (result) {
        return {
          id: result.id,
          resultName: result.resultName,
          link: result.link,
          text: cutString(result.text),
          keywords: result.keywords,
          date: result.date,
        };
      } else {
        return {
          id: null,
          resultName: null,
          link: null,
          text: null,
          keywords: null,
          date: null,
        };
      }
    },
    resultStatistics() {
      if (this.result && this.result.text && this.result.keywords) {
        return this.getTextStatistics(this.result.text, this.result.keywords);
      }
      return null;
    },
  },
  methods: {
    ...mapActions(["fetchResults"]),
    getTextStatistics(text, keywords) {
      const words = text.match(/\w+/g) || [];
      const sentences = text.match(/[^.!?]+[.!?]+/g) || [];
      const keywordCounts = {};

      keywords.forEach((keyword) => {
        const regex = new RegExp(`\\b${keyword}\\b`, "gi");
        const matches = text.match(regex);
        keywordCounts[keyword] = matches ? matches.length : 0;
      });

      return {
        wordCount: words.length,
        characterCount: text.length,
        sentenceCount: sentences.length,
        keywordCounts: keywordCounts,
      };
    },
    showTranslateModal() {
      this.$bvModal.show("translationModal");
    },
    async summarizeText() {
      const apiKey = process.env.VUE_APP_GEMINI_API_KEY;
      try {
        const response = await summarizeTextGemini(apiKey, this.result.text);
        this.summarizedText = response.data.candidates[0].content.parts[0].text;
        this.$refs.summarizationModal.showModal();
      } catch (error) {
        console.error("Error generating summary: ", error);
      }
    },
    async analyzeSentimentText() {
      const apiKey = process.env.VUE_APP_GEMINI_API_KEY;
      try {
        const response = await analyzeSentimentTextGemini(
          apiKey,
          this.result.text
        );
        this.analyzedSentimentText =
          response.data.candidates[0].content.parts[0].text;
        this.$refs.sentimentAnalysisModal.showModal();
      } catch (error) {
        console.error("Error generating sentiment analysis: ", error);
      }
    },
    async translateText() {
      const apiKey = process.env.VUE_APP_GEMINI_API_KEY;
      try {
        const response = await translateTextGemini(
          apiKey,
          this.language,
          this.result.text
        );
        this.translatedText = response.data.candidates[0].content.parts[0].text;
        this.showModal("Translated Text Preview", this.translatedText);
        this.$bvModal.hide("translationModal");
      } catch (error) {
        console.error("Error translating text: ", error);
      }
    },
    canceltranslationModal() {
      this.translationModal = false;
    },
    downloadSummary() {
      downloadFile("summary", this.summarizedText);
    },
    downloadSentimentAnalysis() {
      downloadFile("sentiment", this.analyzedSentimentText);
    },
    cancelPreview() {
      this.$refs.modalComponent.modalVisible = false;
    },
  },
  data() {
    return {
      summarizedText: "",
      summarizedTextTitle: "Summarized Text Preview",
      noSummarizedTextMessage: "No summarized text is available",
      analyzedSentimentText: "",
      analyzedSentimentTextTitle: "Sentiment Analysis Preview",
      noAnalyzedSentimentTextMessage: "No sentiment analysis is available",
      // translatedText: null,
      language: "",
    };
  },
  created() {
    this.fetchResults();
  },
};
</script>

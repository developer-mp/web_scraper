<template>
  <div>
    <NavBarComponent />
    <div class="container-results">
      <h2 class="details-header">Scraping Result Details</h2>
      <ul class="details-list">
        <li class="results-action">
          <b-button-group>
            <b-dropdown right text="Action">
              <b-dropdown-item @click="summarizeText"
                >Summarization</b-dropdown-item
              >
              <b-dropdown-item @click="analyzeSentiment"
                >Sentiment Analysis</b-dropdown-item
              >
              <b-dropdown-item @click="translateText"
                >Translation</b-dropdown-item
              >
            </b-dropdown>
          </b-button-group>
        </li>
        <li><strong>Id:</strong> {{ result.id }}</li>
        <li><strong>Result Name:</strong> {{ result.resultName }}</li>
        <li><strong>Link:</strong> {{ result.link }}</li>
        <li><strong>Keywords:</strong> {{ result.keywords }}</li>
        <li><strong>Text Sample:</strong> {{ result.text }}</li>
        <li><strong>Date:</strong> {{ result.date }}</li>
      </ul>
    </div>
    <FooterComponent />
  </div>
</template>

<style>
.container-results {
  display: flex;
  flex-direction: column;
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

.results-action {
  margin-bottom: 1em;
}
</style>

<script>
import { mapGetters, mapActions } from "vuex";
import FooterComponent from "./../components/FooterComponent.vue";
import NavBarComponent from "./../components/NavBarComponent.vue";
import { cutString } from "./../utils/cutString.js";
import { generateContent } from "./../gemini/geminiApi";

export default {
  components: {
    FooterComponent,
    NavBarComponent,
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
  },
  methods: {
    ...mapActions(["fetchResults"]),
    async summarizeText() {
      const apiKey = process.env.VUE_APP_GEMINI_API_KEY;
      try {
        const response = await generateContent(apiKey, this.result.text);
        this.summarizedText = response.data;
        console.log(this.summarizedText);
      } catch (error) {
        console.error("Error generating content:", error);
      }
    },
  },
  data() {
    return {
      summarizedText: null,
    };
  },
  created() {
    this.fetchResults();
  },
};
</script>

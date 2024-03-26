<template>
  <div>
    <NavBarElement />
    <div class="container-main">
      <h2 class="main-header">Web Scraper</h2>
      <div class="form-wrapper">
        <b-form-input
          v-model="url"
          type="text"
          placeholder="Paste a link"
          class="form-input"
        ></b-form-input>
        <b-form-tags
          v-model="keywords"
          type="text"
          placeholder="Add a keyword"
        ></b-form-tags>
        <div class="button-container">
          <b-button
            @click="submitLink"
            variant="primary"
            class="button"
            :disabled="isSubmitDisabled"
            >Submit</b-button
          >
          <b-button @click="clearFields" variant="secondary" class="button"
            >Clear</b-button
          >
        </div>
      </div>
      <b-modal
        v-model="showPreviewModal"
        id="previewModal"
        title="Preview Results"
        hide-footer
      >
        <div class="preview-results">
          <div v-if="!showPreviewResults">
            <div>The scraping process has been completed successfully.</div>
            <div>Do you want to preview the results?</div>
          </div>
          <div
            v-if="showPreviewResults"
            style="text-align: justify; text-justify: auto"
          >
            <p v-for="(sentence, index) in sentences" :key="index">
              {{ sentence }}
            </p>
          </div>
        </div>
        <div class="button-container">
          <b-button
            v-if="!showPreviewResults"
            class="button mt-3"
            variant="primary"
            @click="previewResults"
            >Preview</b-button
          >
          <b-button
            v-if="showPreviewResults"
            class="button mt-3"
            variant="primary"
            @click="saveResults"
            >Save</b-button
          >
          <b-button
            class="button mt-3"
            variant="secondary"
            @click="cancelPreview"
            >Cancel</b-button
          >
        </div>
      </b-modal>
      <b-modal v-model="showSaveModal" title="Save Results" hide-footer>
        <b-form @submit.prevent="saveResults">
          <b-form-group
            id="resultName"
            label="Result Name"
            label-for="resultNameInput"
          >
            <b-form-input
              id="resultNameInput"
              v-model="resultName"
              type="text"
              required
            ></b-form-input>
          </b-form-group>
          <div class="button-container">
            <b-button
              type="submit"
              variant="primary"
              @click="doneResults"
              class="button mt-3"
              >Done</b-button
            >
            <b-button
              variant="secondary"
              @click="cancelSave"
              class="button mt-3"
              >Cancel</b-button
            >
          </div>
        </b-form>
      </b-modal>
      <FooterElement />
    </div>
  </div>
</template>

<style>
.container-main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.main-header {
  margin-top: 1em;
  margin-bottom: 1em;
}

.form-wrapper {
  width: 40%;
}

.form-input {
  margin-bottom: 1em;
}

.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 1em;
}

.button {
  width: 5em;
}

.preview-results {
  min-height: 7em;
  min-width: 10em !important;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>

<script>
import axios from "axios";
import FooterElement from "./../components/FooterComponent.vue";
import NavBarElement from "./../components/NavBarComponent.vue";
import router from "./../router";

export default {
  components: {
    FooterElement,
    NavBarElement,
  },
  data() {
    return {
      url: "",
      keywords: [],
      showPreviewModal: false,
      showPreviewResults: false,
      sentences: [],
      showSaveModal: false,
      resultName: "",
    };
  },
  computed: {
    isSubmitDisabled() {
      return !this.url.trim() || this.keywords.length == 0;
    },
  },
  methods: {
    async submitLink() {
      try {
        const response = await axios.post("http://localhost:8080/api/scrape", {
          url: this.url,
          keywords: this.keywords,
        });
        this.sentences = response.data.success;
        this.showPreviewModal = true;
      } catch (error) {
        console.error("Error:", error);
      }
    },
    clearFields() {
      this.url = "";
      this.keywords = [];
    },
    previewResults() {
      this.showPreviewResults = true;
    },
    cancelPreview() {
      this.showPreviewResults = false;
      this.$bvModal.hide("previewModal");
    },
    cancelSave() {
      this.showSaveModal = false;
    },
    saveResults() {
      this.showSaveModal = true;
    },
    async doneResults() {
      try {
        await axios.post("http://localhost:8080/api/results", {
          url: this.url,
          keywords: this.keywords,
          resultName: this.resultName,
          sentences: this.sentences,
        });
        router.push("/results");
      } catch (error) {
        console.error("Error:", error);
      }
    },
  },
};
</script>

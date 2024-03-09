<template>
  <div class="container">
    <h2 class="header">Web Scraper</h2>
    <div class="form-wrapper">
      <div class="form-input">
        <b-form-input
          v-model="url"
          type="text"
          placeholder="Paste a link"
        ></b-form-input>
      </div>
      <div>
        <b-form-tags
          v-model="keywords"
          type="text"
          placeholder="Add a keyword"
        ></b-form-tags>
      </div>
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
      title="Preview Result"
      @hidden="cancelPreview"
      hide-footer
    >
      <div class="preview">
        <div v-if="!showPreviewResult">
          <div>The scraping process has been completed successfully.</div>
          <div>Do you want to preview the result?</div>
        </div>
        <div
          v-if="showPreviewResult"
          style="text-align: justify; text-justify: auto"
        >
          <p v-for="(sentence, index) in sentences" :key="index">
            {{ sentence }}
          </p>
        </div>
      </div>
      <div class="preview-button">
        <b-button class="button mt-3" variant="primary" @click="previewResult"
          >Preview</b-button
        >
        <b-button class="button mt-3" variant="secondary" @click="cancelPreview"
          >Cancel</b-button
        >
      </div>
    </b-modal>
  </div>
</template>

<style>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.header {
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

.preview {
  min-height: 7em;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.preview-button {
  display: flex;
  justify-content: space-between;
}
</style>

<script>
import axios from "axios";

export default {
  data() {
    return {
      url: "",
      keywords: [],
      showPreviewModal: false,
      showPreviewResult: false,
      sentences: [],
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
        const response = await axios.post("http://localhost:8080/scrape", {
          url: this.url,
          keywords: this.keywords,
        });
        this.sentences = response.data;
        this.showPreviewModal = true;
      } catch (error) {
        console.error("Error:", error);
      }
    },
    clearFields() {
      this.url = "";
      this.keywords = [];
    },
    previewResult() {
      this.showPreviewResult = true;
    },
    cancelPreview() {
      this.showPreviewResult = false;
      this.$bvModal.hide("previewModal");
    },
  },
};
</script>

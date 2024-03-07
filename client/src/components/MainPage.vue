<template>
  <div
    style="
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
    "
  >
    <h2 style="margin-bottom: 1em">Web Scraper</h2>
    <div style="width: 25%">
      <div style="margin-bottom: 1em">
        <b-form-input
          v-model="url"
          type="text"
          placeholder="Paste a link"
          class="input-focus"
        ></b-form-input>
      </div>
      <div>
        <b-form-tags
          input-id="tags-basic"
          v-model="keywords"
          type="text"
          placeholder="Add a keyword"
        ></b-form-tags>
      </div>
      <div
        style="display: flex; justify-content: space-between; margin-top: 1em"
      >
        <b-button
          @click="submitLink"
          variant="primary"
          style="width: 5em"
          :disabled="isSubmitDisabled"
          >Submit</b-button
        >
        <b-button @click="clearFields" variant="secondary" style="width: 5em"
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
      <div
        class="d-block text-center"
        style="margin-top: 2em"
        v-if="!showPreviewResult"
      >
        <div>The scraping process has been completed successfully.</div>
        <div>Do you want to preview the result?</div>
      </div>
      <div v-if="showPreviewResult">
        <p v-for="(sentence, index) in sentences" :key="index">
          {{ sentence }}
        </p>
      </div>
      <div
        style="display: flex; justify-content: space-between; margin-top: 2em"
      >
        <b-button
          class="mt-3"
          style="width: 5em"
          variant="primary"
          @click="previewResult"
          >Preview</b-button
        >
        <b-button
          class="mt-3"
          style="width: 5em"
          variant="secondary"
          @click="cancelPreview"
          >Cancel</b-button
        >
      </div>
    </b-modal>
  </div>
</template>

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

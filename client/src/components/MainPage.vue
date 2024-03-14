<template>
  <div>
    <NavBarElement />
    <div class="container-main">
      <h2 class="header">Web Scraper</h2>
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
          <b-button
            v-if="!showPreviewResult"
            class="button mt-3"
            variant="primary"
            @click="previewResult"
            >Preview</b-button
          >
          <b-button
            v-else
            class="button mt-3"
            variant="primary"
            @click="doneResult"
            >Done</b-button
          >
          <b-button
            class="button mt-3"
            variant="secondary"
            @click="cancelPreview"
            >Cancel</b-button
          >
        </div>
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

.header {
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
import { saveAs } from "file-saver";
import FooterElement from "./FooterElement.vue";
import NavBarElement from "./NavBarElement.vue";

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
      showPreviewResult: false,
      sentences: [],
      searches: [],
      searchCount: 1,
      fields: [
        { key: "name", label: "Name" },
        { key: "description", label: "Description" },
        { key: "url", label: "URL" },
        { key: "actions", label: "Actions" },
      ],
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
    previewResult() {
      this.showPreviewResult = true;
    },
    saveResult() {
      const textToSave = this.sentences.join("\n");
      const blob = new Blob([textToSave], {
        type: "text/plain;charset=utf-8",
      });
      saveAs(blob, "scraped_data.txt");
      this.cancelPreview();
    },
    doneResult() {
      this.searches.push({
        sentences: this.sentences.slice(),
        url: this.url,
      });
      this.cancelPreview();
      this.searchCount++;
    },
    cancelPreview() {
      this.showPreviewResult = false;
      this.$bvModal.hide("previewModal");
    },
  },
};
</script>
./NavBarElement.vue

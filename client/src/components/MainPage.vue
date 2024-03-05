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
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      url: "",
      keywords: [],
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
        console.log(response.data);
      } catch (error) {
        console.error("Error:", error);
      }
    },
    clearFields() {
      this.url = "";
      this.keywords = [];
    },
  },
};
</script>

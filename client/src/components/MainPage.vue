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
          v-model="text"
          placeholder="Paste a link"
          class="input-focus"
        ></b-form-input>
      </div>
      <div>
        <b-form-tags
          input-id="tags-basic"
          v-model="value"
          placeholder="Add a keyword"
        ></b-form-tags>
      </div>
      <div
        style="display: flex; justify-content: space-between; margin-top: 1em"
      >
        <b-button @click="submitLink" variant="primary" style="width: 5em"
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
export default {
  data() {
    return {
      text: "",
      value: [],
    };
  },
  methods: {
    handlePaste(event) {
      event.preventDefault();
      const pastedText = event.clipboardData.getData("text");
      this.link = pastedText;
    },
    submitLink() {
      console.log("Submitted link:", this.link);
    },
    clearFields() {
      this.text = "";
      this.value = [];
    },
  },
};
</script>

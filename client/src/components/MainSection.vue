<template>
  <div>
    <h2>Web Scraper</h2>
    <div style="display: flex; justify-content: center; align-items: center">
      <input
        v-model="link"
        @paste="handlePaste"
        placeholder="Paste a link here"
        style="
          width: 20%;
          padding: 0.8em;
          border-radius: 0.5em;
          border-width: 0.12em;
          border-color: #35495e;
        "
      />
      <button
        @click="submitLink"
        style="
          padding: 0.8em;
          margin-left: 0.5em;
          border-radius: 0.5em;
          border-width: 0.12em;
          border-color: #35495e;
          background-color: #41b883;
          color: white;
        "
      >
        Submit
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      link: "",
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
  },
};
</script>

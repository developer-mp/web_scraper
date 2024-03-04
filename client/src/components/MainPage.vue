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
    <div style="width: 30%">
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
    </div>
  </div>
</template>

<script>
// import FormTag from "./FormTag.vue";

export default {
  // components: {
  //   FormTag,
  // },
  // data() {
  //   return {
  //     link: "",
  //   };
  // },
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

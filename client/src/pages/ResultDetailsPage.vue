<template>
  <div>
    <NavBarComponent />
    <div class="container-results">
      <h2 class="details-header">Scraping Result Details</h2>
      <li><strong>Id:</strong> {{ result.id }}</li>
      <li><strong>Result Name:</strong> {{ result.resultName }}</li>
      <li><strong>Link:</strong> {{ result.link }}</li>
      <li><strong>Keywords:</strong> {{ result.keywords }}</li>
      <li><strong>Text:</strong> {{ result.text }}</li>
      <li><strong>Date:</strong> {{ result.date }}</li>
      <FooterComponent />
    </div>
  </div>
</template>

<style>
.details-header {
  margin-top: 1em;
  margin-bottom: 1em;
}
</style>

<script>
import FooterComponent from "./../components/FooterComponent.vue";
import NavBarComponent from "./../components/NavBarComponent.vue";

export default {
  components: {
    FooterComponent,
    NavBarComponent,
  },
  props: {
    id: {
      type: Number,
      required: true,
    },
    resultName: {
      type: String,
      required: true,
    },
    link: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
    keywords: {
      type: Array,
      required: true,
    },
    date: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      result: {
        id: null,
        resultName: null,
        link: null,
        text: null,
        keywords: null,
        date: null,
      },
    };
  },
  created() {
    this.result.id = this.id;
    this.result.resultName = this.resultName;
    this.result.link = this.link;
    this.result.text = this.text;
    this.result.keywords = this.keywords;
    this.result.date = this.date;
  },
};
</script>

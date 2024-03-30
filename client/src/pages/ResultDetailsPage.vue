<template>
  <div>
    <NavBarComponent />
    <div class="container-results">
      <h2 class="details-header">Scraping Result Details</h2>
      <ul class="details-list">
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
</style>

<script>
import FooterComponent from "./../components/FooterComponent.vue";
import NavBarComponent from "./../components/NavBarComponent.vue";
import { cutString } from "./../utils/cutString.js";

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
    this.result.text = cutString(this.text);
    this.result.keywords = this.keywords;
    this.result.date = this.date;
  },
};
</script>

e Copy code
<template>
  <div>
    <NavBarElement />
    <div class="container-results">
      <h2 class="results-header">Scraping Results</h2>
      <div>
        <b-table
          striped
          hover
          :items="items"
          :fields="fields"
          @row-clicked="clickRow"
        >
          <template v-slot:cell(actions)="data">
            <b-button @click="handleButtonClick(data.item)" variant="danger"
              >Delete</b-button
            >
          </template>
        </b-table>
      </div>
      <FooterElement />
    </div>
  </div>
</template>

<style>
.container-results {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.results-header {
  margin-top: 1em;
  margin-bottom: 1em;
}
</style>

<script>
import axios from "axios";
import FooterElement from "./../components/FooterComponent.vue";
import NavBarElement from "./../components/NavBarComponent.vue";
import { parseText } from "./../utils/parseText.js";
import { formatDate } from "./../utils/formatDate.js";

export default {
  components: {
    FooterElement,
    NavBarElement,
  },
  data() {
    return {
      items: [],
      fields: [
        { key: "id", label: "Id" },
        { key: "text", label: "Text", formatter: "truncateText" },
        { key: "link", label: "Link", formatter: "truncateLink" },
        { key: "keywords", label: "Keywords" },
        { key: "date", label: "Date" },
        { key: "actions", label: "Actions" },
      ],
    };
  },
  mounted() {
    this.fetchResults();
  },
  methods: {
    async fetchResults() {
      try {
        const response = await axios.get("http://localhost:8080/api/results");
        this.items = response.data.map((item, index) => ({
          id: index + 1,
          text: parseText(item.text),
          link: item.link,
          keywords: item.keywords,
          date: formatDate(item.timestamp),
        }));
      } catch (error) {
        console.error("Error fetching results:", error);
      }
    },
    clickRow(item) {
      // const url = `/results/${item.id}`;
      // this.$router.push(url);
      this.$router.push({ name: "ResultDetails", params: { id: item.id } });
    },
    deleteItem(item) {
      const index = this.items.indexOf(item);
      if (index !== -1) {
        this.items.splice(index, 1);
      }
    },
    truncateText(value) {
      if (value && value.length > 50) {
        return value.substring(0, 50) + "...";
      }
      return value;
    },
    truncateLink(value) {
      if (value && value.length > 30) {
        return value.substring(0, 30) + "...";
      }
      return value;
    },
  },
};
</script>

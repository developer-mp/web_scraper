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
          class="results-column"
        >
          <template v-slot:cell(attachments)="">
            <div class="attachment-icons">
              <i class="fa-solid fa-file" v-b-tooltip title="Original"></i>
              <i class="fa-solid fa-list" v-b-tooltip title="Summary"></i>
              <i
                class="fa-solid fa-chart-simple"
                v-b-tooltip
                title="Analysis"
              ></i>
              <i
                class="fa-solid fa-language"
                v-b-tooltip
                title="Translation"
              ></i>
            </div>
          </template>
          <template v-slot:cell(actions)="data">
            <b-button
              @click="deleteResult(data.item)"
              variant="danger"
              size="sm"
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

.results-column {
  width: 12em;
}

.results-column tbody tr:hover {
  cursor: pointer;
}

.attachment-icons i:not(:last-child) {
  margin-right: 1em;
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
        { key: "resultName", label: "Result Name", thClass: "results-column" },
        { key: "date", label: "Date", thClass: "results-column" },
        { key: "attachments", label: "Attachments" },
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
        this.items = response.data?.map((item, index) => ({
          resultId: item.result_id,
          id: index + 1,
          text: parseText(item.text),
          resultName: item.result_name,
          link: item.link,
          keywords: item.keywords,
          date: formatDate(item.timestamp),
        }));
      } catch (error) {
        console.error("Error fetching results:", error);
      }
    },
    clickRow(item) {
      this.$router.push({ name: "ResultDetails", params: { id: item.id } });
    },
    async deleteResult(item) {
      try {
        await axios.delete(
          `http://localhost:8080/api/results/${item.resultId}`
        );
        window.location.reload();
      } catch (error) {
        console.error("Error deleting result:", error);
      }
    },
  },
};
</script>

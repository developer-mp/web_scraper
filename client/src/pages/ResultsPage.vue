<template>
  <div>
    <NavBarComponent />
    <div class="container-results">
      <h2 class="results-header">Scraping Results</h2>
      <div v-if="results && results.length > 0">
        <b-table
          striped
          hover
          :items="results"
          :fields="fields"
          @row-clicked="clickRow"
          class="results-column"
        >
          <template v-slot:cell(attachments)="">
            <div class="attachment-icons">
              <i class="fa-solid fa-file" v-b-tooltip title="Download"></i>
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
      <div v-else class="no-results">
        <p>No results to display</p>
      </div>
    </div>
    <FooterComponent />
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

.no-results {
  text-align: center;
  margin-top: 20px;
}
</style>

<script>
import { mapGetters, mapActions } from "vuex";
import FooterComponent from "./../components/FooterComponent.vue";
import NavBarComponent from "./../components/NavBarComponent.vue";

export default {
  components: {
    FooterComponent,
    NavBarComponent,
  },
  computed: {
    ...mapGetters(["getResults"]),
    results() {
      return this.getResults;
    },
  },
  data() {
    return {
      fields: [
        { key: "id", label: "Id" },
        { key: "resultName", label: "Result Name", thClass: "results-column" },
        { key: "date", label: "Date", thClass: "results-column" },
        { key: "attachments", label: "Attachments" },
        { key: "actions", label: "Actions" },
      ],
    };
  },
  methods: {
    ...mapActions(["fetchResults"]),
    async deleteResult(item) {
      try {
        await this.$store.dispatch("deleteResult", item.resultId);
      } catch (error) {
        console.error("Error clicking Delete button:", error);
      }
    },
    clickRow(item) {
      this.$router.push({
        name: "ResultDetails",
        params: {
          id: item.id,
        },
      });
    },
  },
  created() {
    this.fetchResults();
  },
};
</script>

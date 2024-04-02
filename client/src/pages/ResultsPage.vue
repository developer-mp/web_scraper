<template>
  <div>
    <NavBarComponent />
    <div class="container-results">
      <h2 class="results-header">Scraping Results</h2>
      <div>
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
    ...mapActions(["fetchResults", "deleteResult"]),
    clickRow(item) {
      this.$router.push({
        name: "ResultDetails",
        params: {
          id: item.id,
          resultName: item.resultName,
          link: item.link,
          text: item.text,
          keywords: item.keywords,
          date: item.date,
        },
      });
    },
  },
  created() {
    this.fetchResults();
  },
};
</script>

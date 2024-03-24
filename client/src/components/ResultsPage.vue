<template>
  <div>
    <NavBarElement />
    <div class="container-results">
      <h2 class="results-header">Scraping Results</h2>
      <div>
        <b-table striped hover :items="items">
          <template #cell(Actions)="data">
            <b-button-group>
              <b-dropdown right split text="Action">
                <b-dropdown-item @click="summarizeItem(data.item)"
                  >Summarize</b-dropdown-item
                >
                <b-dropdown-item @click="sentimentizeItem(data.item)"
                  >Sentimentize</b-dropdown-item
                >
                <b-dropdown-item @click="translateItem(data.item)"
                  >Translate</b-dropdown-item
                >
              </b-dropdown>
            </b-button-group>
          </template></b-table
        >
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
import FooterElement from "./FooterElement.vue";
import NavBarElement from "./NavBarElement.vue";

export default {
  components: {
    FooterElement,
    NavBarElement,
  },
  data() {
    return {
      items: [
        {
          Search: "Search1",
          Link: "google.com",
          Date: "2024-05-09",
          Actions: "Actions",
        },
      ],
    };
  },
  methods: {
    deleteItem(item) {
      const index = this.items.indexOf(item);
      if (index !== -1) {
        this.items.splice(index, 1);
      }
    },
  },
};
</script>

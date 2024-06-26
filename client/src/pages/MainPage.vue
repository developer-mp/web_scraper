<template>
  <div>
    <NavBarComponent />
    <div class="container-main">
      <h2 class="main-header">Web Scraper</h2>
      <div class="form-wrapper">
        <b-form-input
          v-model="url"
          type="text"
          placeholder="Paste a link"
          class="form-input"
        ></b-form-input>
        <b-form-tags
          v-model="keywords"
          type="text"
          placeholder="Add a keyword"
        ></b-form-tags>
        <ButtonGroupComponent
          :is-disabled="isSubmitDisabled"
          confirm-button-label="Submit"
          cancel-button-label="Clear"
          @confirm="submitLink"
          @cancel="clearFields"
        />
      </div>
      <ModalWindowComponent
        ref="scrapingPreviewResultsModal"
        :modalTitle="scrapingPreviewResultsTitle"
        :modalText="modalText"
        @confirm="showSaveResultsModal"
        @cancel="cancelModal"
      />
      <ModalInputComponent
        ref="scrapingSaveResultsModal"
        :modalTitle="scrapingSaveResultsTitle"
        :inputLabel="scrapingSaveResultsLabel"
        :inputValue="resultName"
        :inputPlaceholder="scrapingSaveResultsPlaceholder"
        @confirm="saveScrapingResults"
        @cancel="cancelModal"
        @input="handleResultName"
      />
    </div>
    <FooterComponent />
  </div>
</template>

<script>
import { mapActions } from "vuex";
import FooterComponent from "./../components/FooterComponent.vue";
import NavBarComponent from "./../components/NavBarComponent.vue";
import ModalWindowComponent from "./../components/ModalWindowComponent.vue";
import ModalInputComponent from "./../components/ModalInputComponent.vue";
import ButtonGroupComponent from "./../components/ButtonGroupComponent.vue";
import router from "./../router";

export default {
  components: {
    FooterComponent,
    NavBarComponent,
    ModalWindowComponent,
    ModalInputComponent,
    ButtonGroupComponent,
  },
  data() {
    return {
      url: "",
      keywords: [],
      sentences: [],
      resultName: "",
      scrapingPreviewResultsTitle: "Scraping Results Preview",
      // noScrapingPreviewResultsMessage:
      //   "No sentences found for the given keywords",
      scrapingSaveResultsTitle: "Save Scraping Results",
      scrapingSaveResultsLabel: "Enter Result Name:",
      scrapingSaveResultsPlaceholder: "",
    };
  },
  computed: {
    isSubmitDisabled() {
      return !this.url.trim() || this.keywords.length == 0;
    },
    modalText() {
      if (this.sentences && this.sentences.length > 0) {
        return this.sentences.slice(0, 2).join(" ") + " ...";
      } else {
        return "No sentences found for the given keywords";
      }
    },
  },
  methods: {
    ...mapActions(["submitLink", "saveResults"]),
    async submitLink() {
      try {
        const data = await this.$store.dispatch("submitLink", {
          url: this.url,
          keywords: this.keywords,
        });
        this.sentences = data;
        this.$refs.scrapingPreviewResultsModal.showModal();
      } catch (error) {
        console.error("Error submitting scraping results: ", error);
      }
    },
    clearFields() {
      this.url = "";
      this.keywords = [];
    },
    showSaveResultsModal() {
      this.$refs.scrapingSaveResultsModal.showModal();
    },
    async saveScrapingResults() {
      try {
        await this.$store.dispatch("saveResults", {
          url: this.url,
          keywords: this.keywords,
          resultName: this.resultName,
          sentences: this.sentences,
        });
        router.push("/results");
      } catch (error) {
        console.error("Error saving scraping results: ", error);
      }
    },
    cancelModal() {
      if (this.$refs.scrapingPreviewResultsModal) {
        this.$refs.scrapingPreviewResultsModal.modalVisible = false;
      }
      if (this.$refs.scrapingSaveResultsModal) {
        this.$refs.scrapingSaveResultsModal.modalVisible = false;
      }
    },
    handleResultName(value) {
      this.resultName = value;
    },
  },
};
</script>

<style>
.container-main {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.main-header {
  margin-top: 1em;
  margin-bottom: 1em;
}

.form-wrapper {
  width: 40%;
}

.form-input {
  margin-bottom: 1em;
}
</style>

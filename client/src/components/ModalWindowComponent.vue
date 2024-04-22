<template>
  <b-modal :title="modalTitle" v-model="modalVisible" hide-footer>
    <div v-if="modalText" style="text-align: justify; text-justify: auto">
      {{ modalText }}
    </div>
    <div v-else>{{ modalMessage }}</div>
    <div class="button-container">
      <b-button variant="primary" class="button mt-3" @click="confirmOperation"
        >OK</b-button
      >
      <b-button class="button mt-3" variant="secondary" @click="cancelOperation"
        >Cancel</b-button
      >
    </div>
  </b-modal>
</template>

<script>
export default {
  props: {
    modalTitle: String,
    modalText: String,
    modalMessage: String,
  },
  data() {
    return {
      modalVisible: false,
    };
  },
  methods: {
    confirmOperation() {
      this.$emit("confirm");
      this.modalVisible = false;
    },
    cancelOperation() {
      this.$emit("cancel");
      this.modalVisible = false;
    },
    showModal() {
      this.modalVisible = true;
    },
  },
};
</script>

<style>
.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 1em;
}
</style>

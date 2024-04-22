<template>
  <b-modal :title="modalTitle" v-model="modalVisible" hide-footer>
    <b-form>
      <label style="margin-right: 1em">{{ inputLabel }}</label>
      <input
        type="text"
        :value="inputValue"
        @input="updateInputValue"
        :placeholder="inputPlaceholder"
      />
    </b-form>
    <div class="button-container">
      <b-button variant="primary" @click="confirmOperation" class="button mt-3"
        >Done</b-button
      >
      <b-button variant="secondary" @click="cancelOperation" class="button mt-3"
        >Cancel</b-button
      >
    </div>
  </b-modal>
</template>

<script>
export default {
  props: {
    modalTitle: String,
    inputLabel: String,
    inputValue: String,
    inputPlaceholder: String,
  },
  data() {
    return {
      modalVisible: false,
    };
  },
  methods: {
    confirmOperation() {
      this.$emit("confirm", this.inputValue);
      this.modalVisible = false;
    },
    cancelOperation() {
      this.$emit("cancel");
      this.modalVisible = false;
    },
    updateInputValue(event) {
      this.$emit("input", event.target.value);
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

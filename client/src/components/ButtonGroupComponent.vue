<template>
  <div class="button-container">
    <b-button
      @click="handleConfirmClick"
      variant="primary"
      class="button"
      :disabled="isDisabled"
      >{{ confirmButtonLabel }}</b-button
    >
    <b-button @click="handleCancelClick" variant="secondary" class="button">{{
      cancelButtonLabel
    }}</b-button>
  </div>
</template>

<script>
export default {
  props: {
    isDisabled: {
      type: Boolean,
      default: false,
    },
    confirmButtonLabel: {
      type: String,
    },
    cancelButtonLabel: {
      type: String,
    },
  },
  methods: {
    handleConfirmClick() {
      this.$emit("confirm");
    },
    handleCancelClick() {
      this.$emit("cancel");
    },
  },
};
</script>

<style scoped>
.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 1em;
}

.button {
  width: 5em;
}
</style>

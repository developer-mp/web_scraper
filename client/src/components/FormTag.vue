<template>
  <div style="width: 20%">
    <b-form-tags
      input-id="tags-basic"
      v-model="value"
      placeholder="Add a keyword"
    ></b-form-tags>
  </div>
</template>

<script>
export default {
  data() {
    return {
      value: [],
    };
  },
};
</script>

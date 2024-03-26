<template>
  <div>
    <h2>Scraping Result Details</h2>
    <div>
      <h3>{{ result.id }}</h3>
      <p>{{ result.text }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    // Receive the id parameter as a prop
    id: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      // Simulate fetching result data based on id
      result: {},
    };
  },
  created() {
    // Simulate fetching result data based on id
    this.fetchResult(this.id);
  },
  methods: {
    fetchResult(id) {
      // Simulate fetching result data based on id
      this.result = { id, text: `Result with ID ${id}` };
    },
  },
};
</script>

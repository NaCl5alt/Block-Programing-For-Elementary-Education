<style scoped>
</style>

<template>
  <div>
    <b-container fluid>
      <h1>編集</h1>
    </b-container>
  </div>
</template>

<script>
import axios from "axios";
export default {
  methods: {
    window: (onload = function () {
      axios
        .get(
          "http://localhost:8080/admin/question/${this.$route.params['id']}/edit"
        )
        .then((response) => {})
        .catch((err) => {
          console.log(err);
        });
    }),
  },
};
</script>
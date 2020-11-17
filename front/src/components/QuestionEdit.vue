<style scoped>
</style>

<template>
  <div>
    <b-container fluid>
      <h1>編集</h1>
      <h3>問題タイトル: {{ title }}</h3>
      <!-- <b-form></b-form> -->
    </b-container>
  </div>
</template>

<script>
// import axios from "axios";
export default {
  data() {
    return {
      title: "test",
      content: "",
      answer: "",
    };
  },
  methods: {
    window: (onload = function () {
      // axios
      //   .get("http://localhost/question/${this.$route.params['id']}/detail")
      //   .then((response) => {
      //     this.title = response.title;
      //     this.content = response.content;
      //     this.answer = response.answer;
      //   })
      //   .catch((err) => {
      //     console.log(err);
      //   });
    }),
  },
};
</script>
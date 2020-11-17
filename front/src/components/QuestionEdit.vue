<style scoped>
#question_id {
  font-size: 70%;
  color: gray;
}
</style>

<template>
  <div>
    <b-container>
      <h1>編集</h1>
      <h3>
        問題タイトル: {{ title }} <span id="question_id">(ID: {{ id }})</span>
      </h3>
      <!-- <b-form></b-form> -->
    </b-container>
  </div>
</template>

<script>
// import axios from "axios";
export default {
  data() {
    return {
      id: 0,
      title: "test",
      content: "",
      answer: "",
    };
  },
  methods: {
    window: (onload = function () {
      // ここ修正
      // this.id = this.$route.params["id"];
      // console.log(this.$route.params);
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
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

      <hr class="my-4" />

      <h3>
        タイトル: {{ title }} <span id="question_id">(ID: {{ id }})</span>
      </h3>
      <b-form>
        <b-form-group label="タイトル">
          <b-form-input
            type="text"
            :placeholder="this.datas.title"
          ></b-form-input>
        </b-form-group>

        <b-form-group label="問題文">
          <b-form-textarea :placeholder="this.datas.content"></b-form-textarea>
        </b-form-group>

        <!-- <p>{{ hints }}</p> -->

        <!-- <b-form-select label="ヒントの選択" v-model="selectHintId">
          <b-form-select-option
    
          ></b-form-select-option>
        </b-form-select> -->

        <!-- <b-form-group label="ヒント" v-for="hint in this.hints" :key="hint.id">
          <b-form-input type="text" :placeholder="hint"></b-form-input>
        </b-form-group> -->
      </b-form>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  data() {
    return {
      id: 0,
      title: "test",
      content: "",
      answer: "",
      hints: [],
      datas: {
        qid: 1,
        title: "TEST",
        content: "ASADKSJFNJKDNFKJAN",
        hints: [{ hint: "HINT1" }, { hint: "HINT2" }],
      },
      hintCount: 0,
      selectHintId: 0,
    };
  },
  methods: {
    mtFunc() {
      this.id = this.$route.params["id"];
      this.hintCount = this.datas["hints"].length;

      Axios.get("/api/question/$(this.id)").then((res) => {
        for (let i = 0; i < res.data.hints.length; i++) {
          this.hints = this.hints.push({
            id: i,
            data: res.data.hints[i].hint,
          });
        }
      });
    },
  },
  beforeMount() {
    this.mtFunc();
  },
};
</script>
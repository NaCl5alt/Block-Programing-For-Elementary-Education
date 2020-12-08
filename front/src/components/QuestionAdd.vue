<template>
  <div>
    <b-container>
      <h1>問題追加</h1>
      <div>
        <b-row>
          <b-col sm="3">
            <label for="title">タイトル: </label>
          </b-col>
          <b-col sm="9">
            <b-form-input
              v-model="title"
              type="text"
              name="title"
              placeholder="※ 問題の内容を把握しやすいように命名しましょう。"
              required
            ></b-form-input>
          </b-col>
        </b-row>
        <b-row>
          <b-col sm="3">
            <label for="question">問題文: </label>
          </b-col>
          <b-col sm="9">
            <b-form-textarea
              v-model="question"
              name="question"
              placeholder=""
              required
            ></b-form-textarea>
          </b-col>
        </b-row>
        <b-row>
          <b-col sm="3">
            <label for="answer">答え: </label>
          </b-col>
          <b-col sm="9">
            <b-form-input
              v-model="answer"
              name="answer"
              type="text"
              required
            ></b-form-input>
          </b-col>
        </b-row>
        <b-row>
          <b-col sm="3">
            <label for="hint">ヒント</label>
          </b-col>
          <b-col sm="9">
            <!-- ToDoっぽい感じで追加する -->
            <div>
              <b-row>
                <b-col>
                  <b-form-input
                    v-model="hint_val"
                    name="hint"
                    type="text"
                  ></b-form-input>
                </b-col>
                <b-col>
                  <b-button variant="success" v-on:click="addHint">
                    <!-- <b-icon icon="plus"></b-icon> -->
                    +
                  </b-button>
                </b-col>
              </b-row>
            </div>
            <div>
              <b-list-group>
                <b-row v-for="h in hints" v-bind:key="h.id">
                  <b-col sm="8">
                    <b-list-group-item>
                      {{ h.hint }}
                    </b-list-group-item>
                  </b-col>
                  <b-col sm="2">
                    <b-button variant="danger" v-on:click="deleteHint(h.id)">
                      -
                    </b-button>
                  </b-col>
                </b-row>
              </b-list-group>
            </div>
          </b-col>
        </b-row>

        <b-button
          pill
          variant="outline-primary"
          style="float: right"
          v-on:click="addQuestion"
          >問題を登録</b-button
        >
      </div>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  data() {
    return {
      title: "",
      question: "",
      answer: "",
      hints: [],
      hint_val: "",
      hintCount: 0,
    };
  },
  methods: {
    addHint() {
      // ヒントを追加
      this.hintCount = this.hints.length;
      this.hints.push({ id: this.hintCount, hint: this.hint_val });
      this.hintCount = 0;
      this.hint_val = "";
    },
    deleteHint(id) {
      // ヒントを削除 & idの調整
      this.hints.splice(id, 1);
      for (let i = 0; i < this.hints.length; i++) {
        this.hints[i].id = i;
      }
    },
    async addQuestion() {
      // postHints POST用のヒント達
      var postHints = [];
      // hintsの整形
      this.hints.forEach((hint) => {
        postHints.push({ hint: hint.hint });
      });

      await Axios.post("/api/question", {
        title: this.title,
        content: this.question,
        answer: this.answer,
        hints: postHints,
      },{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          switch (res.status) {
            case 200:
              console.log("追加完了");
              this.$router.go(-1);
              break;
            case 400:
              console.log("リクエストが無効");
              this.$router.go(-1);
              break;
            case 401:
              console.log("トークンエラー");
              this.$router.push("/");
              this.$router.go();
              break;
            default:
              break;
          }
        })
        .catch((err) => {
          console.log(err);
        });
      this.$router.go(-1);
    },
  },
};
</script>

<style scoped>
</style>
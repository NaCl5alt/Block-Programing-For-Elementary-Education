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
        タイトル: {{ title }} <span id="question_id">(ID: {{ this.id }})</span>
      </h3>
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
              required
            ></b-form-input>
          </b-col>
        </b-row>
      </div>
      <b-row>
        <b-col sm="3">
          <label for="question">問題文: </label>
        </b-col>
        <b-col sm="9">
          <b-form-textarea
            v-model="content"
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
                  <b-button variant="danger" @click="deleteHint(h.id)">
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
        style="float: right; margin: 10px"
        variant="outline-danger"
        @click="$bvModal.show('modal-1')"
        >この問題を削除</b-button
      >
      <b-button
        pill
        style="float: right; margin: 10px"
        variant="outline-primary"
        @click="edit"
        >変更を適用</b-button
      >

      <b-modal id="modal-1" hide-footer title="警告!!">
        <div class="d-block text-center">
          <p>
            下の「削除」ボタンを押すと<br />この問題は<span
              style="color: red; font-weight: bold"
              >完全に削除</span
            >されます。<br />よろしいですか？
          </p>
        </div>
        <b-button
          class="mt-3"
          block
          variant="outline-secondary"
          @click="$bvModal.hide('modal-1')"
          >取り消し</b-button
        >
        <b-button class="mt-3" block variant="danger" @click="deleteThis"
          >削除</b-button
        >
      </b-modal>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";

export default {
  data() {
    return {
      id: 0,
      title: "",
      content: "",
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
    async mtFunc() {
      this.id = this.$route.params["id"];

      await Axios.get(`/api/question/contents/${this.id}`,{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          this.title = res.data["title"];
          this.content = res.data["content"];
          this.answer = res.data["answer"];
          for (let i = 0; i < res.data.hints.length; i++) {
            this.hints = this.hints.push({
              id: i,
              hint: res.data.hints[i].hint,
            });
          }
        })
        .catch((error) => {
          console.log(error);
          this.title = "[ERROR] タイトルが存在しません。";
          this.content = "[ERROR] 問題が存在しません。";
          this.answer = "[ERROR] 答えが存在しません。";
          this.hints = this.hints.concat([
            { id: 1, hint: "[ERROR] ヒント1が存在しません。" },
            { id: 2, hint: "[ERROR] ヒント2が存在しません。" },
          ]);
        });
      this.hintCount = this.hints.length;
      this.editable_hintCount = this.hints.length;
    },
    async edit() {
      await Axios.post(`/api/question/${this.id}`, {
        title: this.title,
        content: this.content,
        answer: this.answer,
        hints: this.hints,
      })
        .then((res) => {
          switch (res.status) {
            case 201:
              console.log("追加完了");
              this.$router.go(-1);
              break;
            case 400:
              console.log("リクエストが無効");
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
        .catch((error) => {
          console.log(error);
        });
      this.$router.go(-1);
    },
    async deleteThis() {
      console.log("編集");
      this.$bvModal.hide("modal-1");
      await Axios.delete(`/question/${this.id}`,{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          switch (res.status) {
            case 200:
              console.log("削除完了");
              this.$router.go(-1);
              break;
            case 400:
              console.log("リクエスト無効");
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
        .catch((error) => {
          console.log(error);
        });
      this.$router.go(-1);
    },
  },
  beforeMount() {
    this.mtFunc();
  },
};
</script>
<template>
  <div>
    <b-container>
      <h2>問題一覧</h2>
      <div>
        <b-table-simple striped hover>
          <b-head>
            <b-tr>
              <b-th>No.</b-th>
              <b-th>問題名</b-th>
              <b-th>進捗</b-th>
            </b-tr>
          </b-head>
          <b-tbody>
            <b-tr v-for="q in questions" :key="q.id">
              <b-td>{{ q.id }}</b-td>
              <b-td>{{ q.title }}</b-td>
              <b-td>{{ q.progress }}</b-td>
            </b-tr>
          </b-tbody>

          <!-- <b-tfoot>
            <b-tr variant="secondary" class="text-right">
              問題数: <b>{{ max }}</b>
            </b-tr>
          </b-tfoot> -->
        </b-table-simple>
        <!-- <div class="row">
          <div class="col">No.</div>
          <div class="col-8">問題名</div>
        </div>
        <div v-for="question in questions" :key="question.id" class="row">
          <div class="col">
            <i v-if="question.progress" class="far fa-check-square"></i>
            <i v-else class="far fa-square"></i>
            {{ question.id }}
          </div>
          <div class="col-8">
            {{ question.title }}
          </div>
        </div> -->

        <infinite-loading @infinite="infiniteHandler" spinner="spiral">
          <div slot="spinner">ロード中...</div>
          <div slot="no-more">もう登録されている問題がないよ</div>
          <div slot="no-results">登録されている問題がないよ</div>
        </infinite-loading>
      </div>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  name: "Questions",
  data() {
    return {
      questions: [],
      end: 0,
      max: 0,
      pageNumber: 0,
    };
  },
  methods: {
    async infiniteHandler($state) {
      if (this.end === 0) await this.firstQuestion();
      if (this.end >= this.max) {
        $state.complete();
      } else {
        await this.getQuestions();
        $state.loaded();
      }
    },
    async getCount() {
      Axios.get("/question/count")
        .then((res) => {
          this.max = res.data["count"];
        })
        .catch((error) => {
          console.log(error);
          this.max = 20;
        });
    },
    async firstQuestion() {
      this.questions.push({
        id: 1,
        title: "Title",
        progress: false,
      });
      this.end = this.questions.length;
      console.log(this.questions);
    },
    async getQuestions() {
      this.questions = this.questions.concat({
        id: 2,
        title: "Title",
        progress: true,
      });
      this.end = this.questions.length;
      console.log(this.questions);
    },
  },
  beforeMount() {
    this.getCount();
  },
};
</script>
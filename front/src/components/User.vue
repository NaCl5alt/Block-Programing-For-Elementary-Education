<style scoped>
</style>

<template>
  <div>
    <b-container>
      <h2>
        ユーザの進捗状況
        <span style="color: gray; font-size: 50%">(ID: {{ userId }})</span>
      </h2>
      <div>
        <b-table-simple hover>
          <colgroup style="width: 20%"></colgroup>
          <colgroup style="width: 50%"></colgroup>
          <colgroup style="width: 20%"></colgroup>
          <b-thead head-variant="light">
            <b-tr>
              <b-th class="text-center">No.</b-th>
              <b-th class="text-center">問題名</b-th>
              <b-th class="text-center"></b-th>
            </b-tr>
          </b-thead>
          <b-tbody>
            <b-tr v-for="q in questions" :key="q.id">
              <b-td class="text-right">{{ q.id }}</b-td>
              <b-td class="text-right">{{ q.title }}</b-td>
              <b-td>
                <p
                  v-if="q.progress"
                  :to="{ name: 'Question', params: { id: q.id } }"
                >
                  トライ!
                </p>
                <p v-else :to="{ name: 'Question', params: { id: q.id } }">
                  <b-icon icon="star"></b-icon> 完了
                </p>
              </b-td>
            </b-tr>
          </b-tbody>
        </b-table-simple>
        <infinite-loading @infinite="infiniteHandler" spinner="spiral">
          <div slot="spinner">
            <b-icon
              icon="arrow-counterclockwise"
              animation="spin-reverse"
              font-scale="4"
            ></b-icon
            >ロード中...
          </div>
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
  data() {
    return {
      userId: 0,
      data: [],
      questions: [],
      fields: [
        { key: "id", label: "No. " },
        { key: "title", label: "問題名" },
      ],
      end: 0,
      max: 0,
    };
  },
  methods: {
    async mtFunc() {
      this.userId = this.$route.params["id"];

      await Axios.get("/api/user/$(this.id)/progress")
        .then((res) => {
          this.data = res.data;
        })
        .catch((error) => {
          console.log(error);
          var test_data = {
            progress: [1, 2, 3, 4],
          };
          this.data = test_data;
        });
    },
    async infiniteHandler($state) {
      if (this.end === 0) await this.firstQuestion();
      console.log(this.end);
      console.log(this.max);
      if (this.end >= this.max) {
        $state.complete();
      } else {
        await this.getQuestions();
        $state.loaded();
      }
    },
    async getCount() {
      await Axios.get("/api/question/count")
        .then((res) => {
          this.max = res.data["count"];
        })
        .catch((error) => {
          console.log(error);
          this.max = 500;
        });
    },
    async firstQuestion() {
      await Axios.get("/api/question")
        .then((res) => {
          this.questions = this.questions.concat(res.data);
        })
        .catch((error) => {
          console.log(error);
          this.questions = this.questions.concat({
            id: 1,
            title: "[TEST] Title",
            progress: false,
          });
        });
      this.end = this.questions.length;
      console.log(this.questions);
    },
    async getQuestions() {
      var test = false;
      await Axios.get("/api/question/paging", {
        qid: this.end,
      })
        .then((res) => {
          this.questions = this.questions.concat(res.data);
        })
        .catch((error) => {
          console.log(error);
          for (let i = this.end + 1; i <= this.end + 50; i++) {
            if (i <= this.max) {
              test = Boolean(Math.round(Math.random()));
              this.questions = this.questions.concat({
                id: i,
                title: "[TEST] Title",
                progress: test,
              });
            }
          }
        });
      this.end = this.questions.length;
      console.log(this.questions);
    },
  },
  beforeMount() {
    this.getCount();
    this.mtFunc();
  },
};
</script>
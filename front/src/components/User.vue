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
                <p v-if="q.progress" style="font-weight: bold; color: green">
                  <b-icon icon="star"></b-icon> 完了
                </p>
                <p v-else>未完了</p>
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
      data: [], // ユーザの達成済みの問題ID
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

      await Axios.get(`/api/user/progress/${this.userId}`,{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          this.data = res.data["progress"];
        })
        .catch((error) => {
          console.log(error);
          var test_data = {
            progress: [1, 2, 3, 4, 30],
          };
          this.data = test_data["progress"];
        });
    },
    async infiniteHandler($state) {
      if (this.end === 0) await this.firstQuestion();
      // console.log(this.end);
      // console.log(this.max);
      if (this.end >= this.max) {
        $state.complete();
      } else {
        await this.getQuestions();
        $state.loaded();
      }
    },
    async getCount() {
      await Axios.get("/api/question/count",{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          this.max = res.data["count"];
        })
        .catch((error) => {
          console.log(error);
          this.max = 500;
        });
    },
    async firstQuestion() {
      var buf_progress = false;
      await Axios.get("/api/question",{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          buf_progress = this.data.includes(res.data["qid"]);
          this.questions = this.questions.concat({
            id: res.data["qid"],
            title: res.data["title"],
            progress: buf_progress,
          });
        })
        .catch((error) => {
          console.log(error);
          buf_progress = this.data.includes(1);
          this.questions = this.questions.concat({
            id: 1,
            title: "[TEST] Title",
            progress: buf_progress,
          });
        });
      this.end = this.questions.length;
    },
    async getQuestions() {
      var buf_progress = false;
      await Axios.get(`/api/question/paging?qid=${this.end}`,{
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      })
        .then((res) => {
          for(let i=0;i<res.data.length;i++){
            console.log("res.data")
            console.log(res.data[i])
            buf_progress = this.data.includes(res.data[i]["qid"]);
              this.questions.push({
                id: res.data[i]["qid"],
                title: res.data[i]["title"],
                progress: buf_progress,
              });
          }
        })
        .catch((error) => {
          console.log(error);
          for (let i = this.end + 1; i <= this.end + 50; i++) {
            if (i <= this.max) {
              buf_progress = this.data.includes(i);
              this.questions = this.questions.concat({
                id: i,
                title: "[TEST] Title",
                progress: buf_progress,
              });
            }
          }
        });
      this.end = this.questions.length;
    },
  },
  beforeMount() {
    this.getCount();
    this.mtFunc();
  },
};
</script>
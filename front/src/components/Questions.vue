<template>
  <div>
    <b-container>
      <h2>
        問題一覧
        <!-- <span style="font-size: 50%; color: gray"
          >(問題数: {{ end }}/{{ max }})</span
        > -->
        <b-icon
          id="popover-1"
          icon="exclamation-circle-fill"
          variant="secondary"
          font-scale="0.7"
        ></b-icon>
        <b-popover
          target="popover-1"
          triggers="hover"
          placement="right"
          title="Info"
          >問題に<ruby
            ><rb>挑戦</rb><rp>(</rp><rt>ちょうせん</rt><rp>)</rp></ruby
          >しましょう！！</b-popover
        >
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
                <router-link
                  v-if="q.progress"
                  class="btn btn-info"
                  :to="{ name: 'Question', params: { id: q.id } }"
                  >トライ!</router-link
                >
                <router-link
                  v-else
                  class="btn btn-success"
                  :to="{ name: 'Question', params: { id: q.id } }"
                  ><b-icon icon="star"></b-icon> 完了</router-link
                >
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
  name: "Questions",
  data() {
    return {
      questions: [],
      fields: [
        { key: "id", label: "No." },
        { key: "title", label: "問題名" },
      ],
      end: 0,
      max: 0,
    };
  },
  methods: {
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
  },
};
</script>

<style scoped>
</style>
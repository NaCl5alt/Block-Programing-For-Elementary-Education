<style scoped>
</style>

<template>
  <div>
    <b-container>
      <h1>ユーザ一覧</h1>

      <div>
        <b-table-simple>
          <colgroup style="width: 20%"></colgroup>
          <colgroup style="width: 60%"></colgroup>
          <colgroup style="width: 15%"></colgroup>
          <colgroup style="width: 20%"></colgroup>
          <b-thead head-variant="light">
            <b-tr>
              <b-th class="text-center">ID</b-th>
              <b-th class="text-center">進捗</b-th>
              <b-th class="text-center"></b-th>
              <b-th class="text-center"></b-th>
            </b-tr>
          </b-thead>
          <b-tbody>
            <b-tr v-for="user in users" :key="user.uid">
              <b-td class="text-right">
                {{ user.uid }}
              </b-td>
              <b-td class="text-right">
                <b-progress
                  variant="success"
                  :max="questionCount"
                  height="2rem"
                >
                  <b-progress-bar :value="user.count">
                    <span style="color: black; font-weight: bold">{{
                      user.count
                    }}</span></b-progress-bar
                  >
                </b-progress>
              </b-td>
              <b-td class="text-right">
                <router-link
                  class="btn btn-info"
                  :to="{ name: 'User', params: { id: user.id } }"
                >
                  状況確認
                </router-link>
              </b-td>
              <b-td class="text-right"
                >{{ user.count }}/{{ questionCount }}({{
                  ((user.count / questionCount) * 100).toFixed(2)
                }}%)</b-td
              >
            </b-tr>
          </b-tbody>
        </b-table-simple>
      </div>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  data() {
    return {
      users: [],
      questionCount: 0,
    };
  },
  methods: {
    async getQuestionCount() {
      await Axios.get("/api/question/count")
        .then((res) => {
          this.questionCount = res.data["count"];
        })
        .catch((error) => {
          console.log(error);
          this.questionCount = 50;
        });
    },
    async mtFunc() {
      await Axios.get("/api/user/progress")
        .then((res) => {
          this.users = res.data["prog"];
          this.users.forEach((user) => {
            user["count"] = user["progress"].length;
          });
          // for (const key in res.data) {
          //   count = res.data[key]["progress"].length;
          //   this.users = this.users.concat({
          //     id: key,
          //     progress: res.data[key]["progress"],
          //     count: count,
          //   });
          // }
        })
        .catch((error) => {
          console.log(error);
          var buf_users = {
            prog: [
              {
                uid: 1,
                progress: [1, 2, 3, 4, 5],
              },
              {
                uid: 2,
                progress: [1, 2, 3],
              },
            ],
          };
          this.users = buf_users["prog"];
          this.users.forEach((user) => {
            user["count"] = user["progress"].length;
          });
        });
      console.log(this.users[0]);
    },
  },
  beforeMount() {
    this.getQuestionCount();
    this.mtFunc();
  },
};
</script>
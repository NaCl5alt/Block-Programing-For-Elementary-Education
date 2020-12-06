<style scoped>
</style>

<template>
  <div>
    <b-container>
      <h1>ユーザ一覧</h1>
      <!-- <p>開発中</p> -->
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
            <b-tr v-for="user in users" :key="user.id">
              <b-td class="text-right">
                {{ user.id }}
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
                  :to="{ name: 'User', params: { id: this.user.id } }"
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
    async mtFunc() {
      await Axios.get("/api/question/count")
        .then((res) => {
          this.questionCount = res.data["count"];
        })
        .catch((error) => {
          console.log(error);
          this.questionCount = 50;
        });

      await Axios.get("/user/progress")
        .then((res) => {
          this.users = res.data;
        })
        .catch((error) => {
          console.log(error);
          var count = 0;
          var buf_users = {
            1: { progress: [1, 2, 3, 4, 5] },
            2: { progress: [1, 2] },
          };
          for (const key in buf_users) {
            count = buf_users[key]["progress"].length;
            this.users.push({
              id: key,
              progress: buf_users[key]["progress"],
              count: count,
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
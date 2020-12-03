<style scoped>
</style>

<template>
  <div>
    <b-container>
      <h1>ユーザ一覧</h1>
      <p>開発中</p>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  data() {
    return {
      users: [],
    };
  },
  methods: {
    mtFunc() {
      Axios.get("/api/user/progress").then((res) => {
        for (let i = 0; i < res.data.length(); i++) {
          this.users.push({
            id: i,
            data: res.data["uid"],
          });
        }
      });
    },
  },
};
</script>
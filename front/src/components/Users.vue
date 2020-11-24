<style scoped>
</style>

<template>
  <div>
    <b-container>
      <h1>ユーザ一覧</h1>
      <ul>
        <li v-for="md in mainDatas" :key="md.id">{{ md.data }}</li>
      </ul>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  data() {
    return {
      mainDatas: [],
    };
  },
  methods: {
    mtFunc() {
      Axios.get("/api/user/progress").then((res) => {
        var buf = res.data;
        for (let i = 0; i < buf.length(); i++) {
          this.mainDatas = this.mainDatas.push({
            id: i,
            data: res.data["uid"],
          });
        }
      });
    },
  },
};
</script>
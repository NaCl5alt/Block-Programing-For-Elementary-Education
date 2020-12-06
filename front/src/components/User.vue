<style scoped>
</style>

<template>
  <div>
    <b-container>
      <h2>各ユーザの状況</h2>
      <p>{{ this.userId }}</p>
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
        });
    },
  },
  beforeMount() {
    this.mtFunc();
  },
};
</script>
<style scoped>
h1 {
  font-style: italic;
  text-align: center;
}
#body-div {
  padding: 10px;
}
</style>

<template>
  <div id="body-div">
    <b-container fluid>
      <b-row>
        <!-- アプリタイトル -->
        <b-col sm="7">
          <b-jumbotron
            bg-variant="info"
            text-variant="white"
            border-variant="dark"
            header="For Educations!"
            lead="Programming Beginner"
          ></b-jumbotron>

          <!-- <div id="for-develop">
            <p class="h2">開発用</p>
            <b-list-group>
              <b-list-group-item to="/login">ログイン画面へ</b-list-group-item>
              <b-list-group-item to="/regist">新規登録画面へ</b-list-group-item>
              <b-list-group-item to="/question"
                >問題一覧画面へ</b-list-group-item
              >
              <b-list-group-item to="/question/1"
                >問題画面へ(id=1)</b-list-group-item
              >
              <b-list-group-item to="/admin">管理者画面へ</b-list-group-item>
            </b-list-group>
          </div> -->
          <div>
            <p class="h2">About</p>
            <b-list-group>
              <b-list-group-item to="#">このアプリについて</b-list-group-item>
              <b-list-group-item to="#">問い合わせ</b-list-group-item>
            </b-list-group>
          </div>
        </b-col>
        <Login />
      </b-row>
    </b-container>
  </div>
</template>

<script>
import Axios from "axios";
import Login from "./Login.vue";
export default {
  components: {
    Login,
  },
  methods: {
    tokenCheck() {
      Axios.post("/api/token").then((res) => {
        switch (res.status) {
          case 200:
            this.$router.push({ path: "/question" });
            this.$router.go();
            break;

          default:
            this.$cookies.remove("token"); // トークンを削除して
            this.$router.push({ path: "/" });
            this.$router.go(); // リダイレクト
            break;
        }
      });
    },
  },
  beforeMount() {
    this.tokenCheck();
  },
};
</script>
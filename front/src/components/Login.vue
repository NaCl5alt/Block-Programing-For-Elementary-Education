<template>
  <div id="body-div">
    <b-container fluid>
      <b-row>
        <b-col> </b-col>
        <!-- ログインフォーム -->
        <b-col>
          <div id="form-div">
            <b-card bg-variant="light">
              <h2>ログイン</h2>
              <b-form>
                <b-form-group
                  label="ユーザID"
                  label-cols-sm="4"
                  description="未登録の方は新規登録をしてください。">
                  <b-form-input v-model="user_id" type="text" name="user_id"></b-form-input>
                </b-form-group>
                <p v-if="nulluserid" class="miss">ユーザIDを入力してください。</p>
                <b-form-group label="パスワード" label-cols-sm="4">
                  <b-form-input
                    v-model="password"
                    type="password"
                    name="password">
                  </b-form-input>
                </b-form-group>
                <p v-if="nullpassword" class="miss">パスワードを入力してください。</p>
                <p v-if="not_login" class="miss">ログインできませんでした。</p>
                <b-row>
                  <b-col></b-col>
                  <b-col>
                    <b-button pill variant="primary" v-on:click="login">ログイン</b-button>
                  </b-col>
                  <b-col>
                    <b-button pill type="reset" variant="danger">リセット</b-button>
                  </b-col>
                </b-row>
              </b-form>
            </b-card>
          </div>
        </b-col>
        <b-col></b-col>
      </b-row>
    </b-container>
  </div>
</template>

<style>
.card-body {
  width: 391px;
}
.miss {
  color: red;
}
</style>

<script>
import axios from "axios";
export default {
  name: "Login",
  data() {
    return {
      user_id: "",
      password: "",
      not_login: false,
      nulluserid: false,
      nullpassword: false,
    }
  },
  methods: {
    async login() {
      if (this.user_id != "" && this.password != "") {
        await axios.post("/api/login", {
            userid: this.user_id,
            password: this.password,
        }).then((res) => {
          if (res.status === 200) {
            this.not_login = this.nulluserid = this.nullpassword = false;
            this.$cookies.set("token", res.data.token)
            this.$cookies.set("admin", res.data.admin)
            var path = "/question";
            if (this.$route.query.redirect) path = this.$route.query.redirect
            this.$router.push({ path: path });
          }
        }).catch(() => {
          this.not_login = true;
        })
      } else {
        if (this.user_id == "") this.nulluserid = true
        if (this.password == "") this.nullpassword = true
      }
    }
  }
}
</script>
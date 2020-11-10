<template>
    <div>
      <b-container class="justifiy-content-center">
        <section class="ex__box">
          <h5>新規登録</h5>
          <p><b-form-input type="text" v-model="mailAddress" placeholder="ユーザid"/></p>
          <p v-if="error.mailAddress">{{ error.mailAddress }}</p>
          <p><b-form-input type="password" v-model="password" placeholder="パスワード"/></p>
          <p v-if="error.password">{{ error.password }}</p>
          <b-button block variant="primary">新規登録</b-button>
        </section>
      </b-container>
    </div>
</template>

<script>
const regeMailAddress = /^[A-Za-z0-9]{1}[A-Za-z0-9_.-]*@{1}[A-Za-z0-9_.-]{1,}\.[A-Za-z0-9]{1,}$/;
const regePassword = /^(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d)[a-zA-Z\d]{8,100}$/;

export default {
  data: function(){
    return {
        mailAddress: '',
        password: '',
        error: {},
    }
  },
  watch: {
    mail: function(val) {
      if (regeMailAddress.test(val)) {
        this.error.mailAddress = "";
      } else if (val === "") {
        this.error.mailAddress = "ユーザidを入力してください";
      } else {
        this.error.mailAddress = "ユーザidの形式に誤りがあります";
      }
    },
    tell: function(val) {
      if (regePassword.test(val)) {
        this.error.password = "";
      } else if (val === "") {
        this.error.password = "パスワードを入力してください";
      } else {
        this.error.password = "パスワードの形式に誤りがあります";
      }
    },
  },
  methods: {
  },
}
</script>
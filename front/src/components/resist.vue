<template>
    <div>
      <b-container class="justifiy-content-center">
        <section class="ex__box">
          <h5>新規登録</h5>
          <p><b-form-input type="text" v-model="userid" placeholder="ユーザid"/></p>
          <p v-show="erroruserid"><font color="red">ユーザidを入力してください</font></p>
          <p><b-form-input type="password" v-model="password" placeholder="パスワード"/></p>
          <p v-show="errorpassword"><font color="red">パスワードを入力してください</font></p>
          <b-button block variant="primary" @click="click">新規登録</b-button>
        </section>
      </b-container>
    </div>
</template>

<script>
export default {
  data: function(){
    return {
        userid: '',
        password: '',
        erroruserid: false,
        errorpassword: false
    }
  },
  methods: {
    regist () {
      axios.post(process.env.VUE_APP_API_URL_BASE + 'user', {
        userid: this.userid,
        password: this.password
      }).then(res => {
        this.$router.push({ path: '/login' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    },
    click: function () {
      if (this.userid === '') {
        this.erroruserid = true
      }
      if (this.password === '') {
        this.errorpassword = true
      }
      if (this.erroruserid = false, this.errorpassword = false) {
        regist ()
      }
    }
  }
}
</script>

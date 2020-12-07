<template>
    <div>
      <b-container class="justifiy-content-center">
        <section class="ex__box">
          <h5>新規登録</h5>
          <p><b-form-input type="text" v-model="userid" placeholder="ユーザid"/></p>
          <p v-show="erroruserid"><font color="red">ユーザidを入力してください</font></p>
          <p v-show="erroruserid_used"><font color="red">このユーザidは使用されています</font></p>
          <p><b-form-input type="password" v-model="password" placeholder="パスワード"/></p>
          <p v-show="errorpassword"><font color="red">パスワードを入力してください</font></p>
          <b-button block variant="primary" @click="click">新規登録</b-button>
        </section>
      </b-container>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data: function(){
    return {
        userid: '',
        password: '',
        erroruserid: false,
        erroruserid_used: false,
        errorpassword: false
    }
  },
  methods: {
    regist () {
      axios.post('/api/user', {
        userid: this.userid,
        password: this.password
      }).then(res => {
        this.$router.push({ path: '/login' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    },
    check () {
      axios.post('/api/user/check', {
        userid: this.userid,
      }).then(res => {
        console.log(res)
        this.erroruserid_used = res.data.exist
      }).catch(error => {
        console.log(error)
      })
    },
    click () {
      if (this.userid !== '' && this.password !== '') {
        this.check()
        if (this.erroruserid_used === false) {
          this.erroruserid = false
          this.errorpassword = false
          this.regist()
        }
      }
      else if (this.userid === '' && this.password === '') {
        this.erroruserid_used = false
        this.erroruserid = true
        this.errorpassword = true
      }
      else if (this.userid === '') {
        this.erroruserid_used = false
        this.erroruserid = true
        this.errorpassword = false
      }
      else if (this.password === '') {
        this.erroruserid_used = false
        this.erroruserid = false
        this.errorpassword = true
      }
    }
  }
}
</script>

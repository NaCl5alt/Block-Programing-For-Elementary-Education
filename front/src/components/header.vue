<template>

<div id="header">
  <div class="border border-info">
    <div class="bg-info">
    <div class="section1 text-left" style="float: left;">
      <a href="/">
      <img src="../assets/logo.png" height="50" width="50" alt="ロゴ">
      </a>
      </div>

    <div class="section1 text-right" style="float: right;">

      <span v-show='token !== null && admin !== "true"'>
      <b-button squared
      size="lg"
      variant="success"
      style="width:200px; height:50px;"
      href="/user"
      v-if='admin!=="true"'
      >ユーザーページ
      </b-button>
      <b-button squared
      size="lg"
      variant="success"
      style="width:200px; height:50px;"
      href="/admin"
      v-else
      >管理者ページ
      </b-button>
 <b-button squared
      size="lg"
      variant="danger"
      style="width:200px; height:50px;"
      href="/logout"
      >ログアウト
      </b-button>
      </span>
      <span v-show='token!==null && admin === "true" '>
      <b-button squared
      size="lg"
      variant="success"
      style="width:200px; height:50px;"
      href="/admin"
      >管理者ページ
      </b-button>
      <b-button squared
      size="lg"
      variant="danger"
      style="width:200px; height:50px;"
      href="/logout"
      >ログアウト
      </b-button>
      </span>
      <span v-show='token === null && (admin === "false" || admin === null)'>
        <b-button squared
        size="lg"
        variant="success"
        style="width:200px; height:50px;"
        href="/login"
        >ログイン
        </b-button>
        <b-button squared
        size="lg"
        variant="danger"
        style="width:200px; height:50px;"
        href="/regist"
        >新規登録
      </b-button>
      </span>
      </div>
      <div class="section1 clear" style="clear: both;">
      </div>

      </div>
    </div>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
    token:this.$cookies.get('token'),
    admin:this.$cookies.get('admin'),
    }
  },
  methods:{
    verifyfunc () {
      axios.post('/api/token', {}, { 
        withCredentials: true,
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      }).then(res => {
        console.log('status: ' + res.status)
        if (res.status === 200) {
          this.verify = true
        }
      }).catch(err => {
        // console.log(err)
        if(err.response.status === 401){
          this.verify = false
          this.tokenget()
        } else {
          this.verify = false
          this.$cookies.remove('token')
          this.$cookies.remove('admin')
        }
     })
    },
    tokenget(){
      axios.get('/api/token', { withCredentials: true,
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      }).then(res => {
        // console.log('status: ' + res.status)
        if (res.status === 200) {
          this.$cookies.remove('token')
          this.$cookies.remove('admin')
          this.$cookies.set('token', res.data.token)
          this.$cookies.set('admin', res.data.admin)
          this.verify = true
        }
      }).catch(() => {
        // console.log(err)
        this.$cookies.remove('token')
        this.$cookies.remove('admin')
        this.verify = false
      })
   }
  },
  async beforeMount () {
    await this.verifyfunc()
  }
}
</script>

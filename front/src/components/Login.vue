

<template>
  <div id="body-div">
    <b-container fluid>
       <b-row>
        <b-col>
        </b-col>
        <!-- ログインフォーム -->
        <b-col>
          <div id="form-div">
            <b-card bg-variant="light">
              <h2>ログイン</h2>

              <b-form>
                <b-form-group
                  label="ユーザID"
                  label-cols-sm="4"
                  description="未登録の方は新規登録をしてください。"
                >
                  <b-form-input
                    v-model="user_id"
                    type="text"
                    name="user_id"
                  ></b-form-input>
                </b-form-group>
                
                <p v-if=missuserid class="miss">
                ユーザIDが違います。
                </p>
                

                <b-form-group label="パスワード" label-cols-sm="4">
                  <b-form-input
                    v-model="password"
                    type="password"
                    name="password"
                  ></b-form-input>
                </b-form-group>
                
                <p v-if=misspassword class="miss">
                パスワードが違います。
                </p>
                

                <b-row>
                  <b-col></b-col>
                  <b-col>
                    <b-button pill variant="outline-primary" v-on:click="login">ログイン</b-button>
                  </b-col>
                  <b-col>
                    <b-button pill type="reset" variant="outline-danger"
                      >リセット</b-button
                    >
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
.card-body{
  width:391px;
}
.miss{
  color:red;
}
</style>

<script>
import axios from 'axios'
export default {
  name: 'Login',
  data (){
    return{
      user_id:'',
      password:'',
      missuserid: false,
      misspassword: false
    }
  },
  methods:{
    login(){
      if(this.user_id=='') this.missuserid = true
      if(this.password=='') this.misspassword = true
      console.log(this.user_id)
      console.log(this.password)
      if(this.missuserid==false &&this.misspassword==false){
        axios.post('/api/login',{
          user_id: this.user_id,
          password: this.password
        }).then(res => {
          console.log(res.data)
        }).catch(error => {
          console.log(error)
        })
      }
    
    }
  }
}
</script>
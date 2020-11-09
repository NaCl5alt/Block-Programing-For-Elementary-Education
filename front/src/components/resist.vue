<template>
    <div>
      <b-container class="justifiy-content-center">
        <section class="ex__box">
          <h5>新規登録</h5>
          <p><b-form-input type="text" v-model="mailAddress" placeholder="メールアドレス"/></p>
          <p><b-form-input type="password" v-model="password" placeholder="パスワード"/></p>
          <b-button block variant="primary">
            <a class="button--green">新規登録</a>
          </b-button>
        </section>
      </b-container>
    </div>
</template>

<script>
export default {
  data: function(){
    return {
        mailAddress: '',
        password: '',
    }
  },
  methods: {
  },
}
</script>
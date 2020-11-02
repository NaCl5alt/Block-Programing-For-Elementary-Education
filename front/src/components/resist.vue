<template>
    <div>
      <section class="ex__box">
        <h5>新規登録</h5>
        <p><input type="text" v-model="mailAddress" placeholder="メールアドレス"></p>
        <p><input type="password" v-model="password" placeholder="パスワード"></p>
        <b-button class="links">
          <a class="button--green">新規登録</a>
        </b-button>
      </section>
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
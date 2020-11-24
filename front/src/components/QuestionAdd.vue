<template>
  <div>
    <h1>問題追加</h1>

    <b-form>
      <b-form-group
        label="タイトル"
        description="問題の内容を把握しやすいように命名しましょう。"
      >
        <b-form-input type="text" required></b-form-input>
      </b-form-group>
        <div id="app-2">
          <input type="text" v-model="message">
          <div class="error" v-if="error.require">必須項目です。</div>
          <div class="error" v-if="error.tooLong">長すぎます。</div>
        </div>


      <b-form-group
        label="問題"
        
      >
        <b-form-input type="text" required></b-form-input>
      </b-form-group>

      <b-form-group
        label="答え"
        description=""
      >
        <b-form-input type="text" required></b-form-input>
      </b-form-group>

      <b-form-group
        label="ヒント"
        description=""
      >
        <b-form-input type="text" required></b-form-input>
      </b-form-group>
        <b-button type="submit" variant="primary">Submit</b-button>
      <b-button type="reset" variant="danger">Reset</b-button>
    </b-form>

    


  </div>

  

</template>

<script>
export default {};
</script>

<style scoped>
</style>
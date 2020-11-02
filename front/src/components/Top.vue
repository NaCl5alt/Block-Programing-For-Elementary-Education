<style scoped>
h1 {
  font-style: italic;
  text-align: center;
}
</style>

<template>
  <div id="body-div">
    <b-container fluid>
      <b-row>
        <!-- アプリタイトル -->
        <b-col sm="7">
          <b-jumbotron
            header="For Educations!"
            lead="Programming Beginner"
          ></b-jumbotron>
          <!-- <h1>For Educations!</h1> -->

          <div>
            <p class="h2">開発用</p>
            <b-list-group>
              <b-list-group-item to="/login">ログイン画面へ</b-list-group-item>
              <b-list-group-item to="/regist">新規登録画面へ</b-list-group-item>
              <b-list-group-item to="/question"
                >問題一覧画面へ</b-list-group-item
              >
              <b-list-group-item to="/admin">管理者画面へ</b-list-group-item>
            </b-list-group>
          </div>
        </b-col>

        <!-- ログインフォーム -->
        <b-col sm="5">
          <div id="form-div">
            <b-card bg-variant="light">
              <h2>ログイン</h2>

              <b-form>
                <b-form-group label="ユーザID" label-cols-sm="4">
                  <b-form-input type="text" name="user_id"></b-form-input>
                </b-form-group>

                <b-form-group label="パスワード" label-cols-sm="4">
                  <b-form-input type="password" name="password"></b-form-input>
                </b-form-group>

                <b-row>
                  <b-col></b-col>
                  <b-col>
                    <b-button pill type="submit" variant="outline-primary"
                      >ログイン</b-button
                    >
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
      </b-row>
    </b-container>
  </div>
</template>

<script>
export default {};
</script>
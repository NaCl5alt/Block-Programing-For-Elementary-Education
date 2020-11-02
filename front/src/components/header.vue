<template>
  <div id="header">
    <div class="section1 text-left" style="float: left;">
      <img src="../assets/logo.png" height="50" width="50" alt="ロゴ" href="/">
      </div>
    <div class="section1 text-right" style="float: right;">
      <span v-if="userlogin==true">
      <b-button squared
      size="lg"
      variant="outline-success"
      style="width:150px; height:50px;"
      href="/user"
      >ユーザーページ
      </b-button>
      <b-button squared
      size="lg"
      variant="outline-danger"
      style="width:150px; height:50px;"
      href="/logout"
      >ログアウト
      </b-button>
      </span>
      <span v-else-if="adminlogin==true">
      <b-button squared
      size="lg"
      variant="outline-success"
      style="width:150px; height:50px;"
      href="/admin"
      >管理者ページ
      </b-button>
      <b-button squared
      size="lg"
      variant="outline-danger"
      style="width:150px; height:50px;"
      href="/logout"
      >ログアウト
      </b-button>
      </span>
      <span v-else>
        <b-button squared
        size="lg"
        variant="success"
        style="width:150px; height:50px;"
        href="/login"
        >ログイン
        </b-button>
        <b-button squared
        size="lg"
        variant="danger"
        style="width:150px; height:50px;"
        href="/regist"
        >新規登録
      </b-button>
      </span>
      </div>
      <div class="section1 clear" style="clear: both;">
      </div>
    </div>
</template>

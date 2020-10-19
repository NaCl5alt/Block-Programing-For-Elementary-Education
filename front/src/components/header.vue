<template>
    <div class="header">
      <h1>logo</h1>
      <span v-if="userlogin==true">
        <button @click="mypage" type="button" value="mypage" >マイページ</button>
        <button @click="logout" type="submit" value="logout" >ログアウト</button>
      </span>
      <span v-else-if="adminlogin==true">
        <button @click="adminpage" type="button" value="adminpage" >管理者ページ</button>
        <button @click="logout" type="submit" value="logout" >ログアウト</button>
      </span>
      <span v-else>
        <button @click="login" type="submit" value="login" >ログイン</button>
        <button @click="register" type="submit" value="register" >新規登録</button>
      </span>
    </div>

</template>

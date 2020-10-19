<template>
    <div class="hello">
      <div id="app-104">
      <button @click="user" type="submit" value="userpage" v-if="userlogin">userpage</button>
      <button @click="question" type="button" value="questionadd" v-else-if="adminlogin">questionadd</button>
      </div>
    </div>
    <script>

    </script>
</template>

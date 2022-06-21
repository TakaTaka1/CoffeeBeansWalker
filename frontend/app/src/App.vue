<template>
  <div id="app">
    <Header/>
    <router-view />
  </div>
</template>

<script>
  // 参考サイト https://blog.narito.ninja/detail/206
  import Header from "./components/Header.vue";

   export default {
        name: 'app',
        components: {Header}
    }

</script>

<style>
@import url('https://use.fontawesome.com/releases/v5.0.10/css/all.css');
@import url('https://fonts.googleapis.com/css2?family=Caveat&display=swap');
#app {
  height: 100%;
  width: 100%;
}
</style>
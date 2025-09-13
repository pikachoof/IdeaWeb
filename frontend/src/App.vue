<script setup lang="ts">
import Header from './components/elements/Header.vue'
import HomePageBody from './components/pages/HomePageBody.vue';
import Footer from './components/elements/Footer.vue';
</script>

<template>
    <div class="flex flex-col min-h-screen">
      <Header />
      <HomePageBody />
      <Footer />
    </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

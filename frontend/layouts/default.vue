<template>
  <div class='min-h-screen flex flex-col '>
    <Navbar />
    <NuxtPage  class="grow "/>
    <Footer  />
  </div>
</template>

<script setup>
import Navbar from "~/components/NavBar.vue";
import Footer from "~/components/Footer.vue";
</script>

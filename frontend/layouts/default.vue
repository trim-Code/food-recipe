<template>
  <div class='min-h-screen flex flex-col  '>
    <Navbar  />
    <NuxtPage  class="flex-grow mt-20  "/>
    <Footer v-if="showNavbar"/>
  </div>
</template>

<script setup>
import Navbar from "~/components/NavBar.vue";
import Footer from "~/components/Footer.vue";

import { useRoute } from 'vue-router';
const route=useRoute();
const showNavbar=route.name !== 'contacts' && route.name !== 'settings'
</script>

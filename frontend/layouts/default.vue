<template>
  <div class='min-h-screen flex flex-col  '>
    <Navbar  />
    <NuxtPage  class="flex-grow mt-20  "/>
    <Footer v-if="showFooter"/>
  </div>
</template>

<script setup>
import Navbar from "~/components/NavBar.vue";
import Footer from "~/components/Footer.vue";

import { useRoute } from 'vue-router';
import { computed } from 'vue';
const route = useRoute();
const showFooter = computed(() => !['settings','contacts'].includes(route.name));
</script>

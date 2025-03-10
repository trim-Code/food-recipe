<template>
  <footer
    class="fixed bottom-0 left-0 w-full bg-gradient-to-b from-[#0b304d] to-[#993b0cc5] text-white"
  >
    <svg
      class="absolute top-0 w-full"
      viewBox="0 0 1440 320"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fill="#000"
        fill-opacity="1"
        d="M0,64L48,80C96,96,192,128,288,122.7C384,117,480,75,576,64C672,53,768,75,864,101.3C960,128,1056,160,1152,181.3C1248,203,1344,213,1392,218.7L1440,224L1440,320L1392,320C1344,320,1248,320,1152,320C1056,320,960,320,864,320C768,320,672,320,576,320C480,320,384,320,288,320C192,320,96,320,48,320L0,320Z"
      ></path>
    </svg>
    <img src="assets/images/2.png" class="absolute w-64" />
    <img src="assets/images/4.png" class="absolute w-64 right-0 -top-16" />
    <div class="container mx-auto px-4 py-12 text-center relative z-10">
      <h2 class="text-3xl font-bold mb-4">Cook Share & Enjoy</h2>
      <p class="text-lg mb-6">
        Join our community and explore endless recipes for every taste and
        occasion.
      </p>
      <div class="flex justify-center space-x-4">
        <NuxtLink to="/" class="text-white hover:text-orange-200"
          >Home</NuxtLink
        >
        <NuxtLink to="/my-recipes" class="text-white hover:text-orange-200"
          >Recipes</NuxtLink
        >
        <NuxtLink to="/contacts" class="text-white hover:text-orange-200"
          >Contact</NuxtLink
        >
      </div>
      <p class="mt-6 text-sm">&copy; 2025 Kimem. All rights reserved.</p>
    </div>
  </footer>
</template>

<style scoped>
footer {
  position: relative;
  bottom: 0;
  width: 100%;
}
</style>

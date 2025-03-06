<template>
  <div class="w-screen">
    <Navbar class="" />

    <NuxtPage />
    <div class="relative h-screen w-screen ">
      <div
        class="absolute inset-0 bg-[url('/assets/images/bg-home.jpg')] bg-cover bg-no-repeat bg-center"
      >
    </div>
      <div class="absolute inset-0 bg-black bg-opacity-40"></div>
      <div class="absolute inset-0 bg-black bg-opacity-50"></div>
      <div class="relative flex flex-col items-center justify-center h-full w-3/4 m-auto" >
        <h1 class="text-4xl w-2/3 text-center font-bold text-[#faebefe5] mb-4">
          Master new flavors and turn every 
            <span class="bg-[#24515b] px-3 inline-block transform -skew-x-12">meal</span>
          into a masterpiece!
        </h1>
        <p class="mt-3 pb-6 text-lg text-center  w-2/3 text-[#faebefc9]">Unlock the secrets of delicious cooking. Explore Cook and Savor Every Bite! Shop  cook and savor – your favorite recipes anytime anywhere!</p>
        
        <button class=" btn relative -bottom-40 z-60  group">
          <span class="">Explore Now</span>
            <font-awesome-icon
          icon="arrow-down" 
          class=" w-4 text-[#302a2b] 
group-hover:size-5
          "
        />

        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
useHead({
  title: "Recipe App",
  meta: [
    {
      name: "description",
      content:
        "A comprehensive, user-friendly app for finding and sharing delicious recipes.",
    },
  ],
});
</script>

<style scoped>

</style>
